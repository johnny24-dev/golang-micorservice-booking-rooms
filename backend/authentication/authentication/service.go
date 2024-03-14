package authentication

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/types"
	"github.com/nekizz/final-project/backend/authentication/constant"
	"github.com/nekizz/final-project/backend/authentication/logger"
	"github.com/nekizz/final-project/backend/authentication/util"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"

	j "github.com/dgrijalva/jwt-go"
	"github.com/nekizz/final-project/backend/authentication/jwt"
	pb "github.com/nekizz/final-project/backend/go-pbtype/authentication"
	pbn "github.com/nekizz/final-project/backend/go-pbtype/notification"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) ChangePassword(ctx context.Context, r *pb.Authentication) (*pbn.Notification, error) {
	if err := validateAuthentication(r); nil != err {
		return nil, err
	}

	err := NewRepository(s.db).UpdateOne(getModel(uint(r.GetId()), prepareDataToRequest(r)))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("change password failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.ChangePasswordFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	return &pbn.Notification{
		Status:  "OK",
		Message: constant.ChangePasswordSuccessfully,
		Code:    http.StatusOK,
	}, nil
}

func (s *Service) Login(ctx context.Context, r *pb.Authentication) (*pb.AuthenticationResponse, error) {
	if err := validateAuthentication(r); nil != err {
		return nil, err
	}

	account, err := NewRepository(s.db).Login(r.GetUsername(), r.GetPassword())
	if err != nil {
		return nil, err
	}
	if account.ID == 0 {
		return nil, errors.New("Wrong username or password!")
	}

	rdb := util.NewRedis(2)
	keyRedis := r.GetUsername() + "_jwt_token"
	value, err := util.GetRedis(rdb, keyRedis)
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("get jwt_token from redis failed")
	}
	if value != nil {
		jwtToken, ok := value.(string)
		if ok {
			return &pb.AuthenticationResponse{
				JwtToken: jwtToken,
			}, nil
		}
	}

	customerId, err := NewRepository(s.db).GetCustomerIdByAccountId(int(account.ID))
	if err != nil {
		return nil, err
	}

	jwtToken, err := jwt.JWTAuthService().GenerateToken(account.Username, customerId, true)
	if err != nil {
		return nil, err
	}

	err = util.SetRedis(rdb, keyRedis, jwtToken, 2*time.Hour)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("set jwt_token to redis failed")
	}

	return &pb.AuthenticationResponse{
		JwtToken: jwtToken,
	}, nil
}

func (s *Service) SignUp(ctx context.Context, r *pb.Authentication) (*pb.Authentication, error) {
	if err := validateAuthentication(r); nil != err {
		return nil, err
	}

	account, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("create account failed")
		return nil, err
	}

	return prepareDataToResponse(account), nil
}

func (s *Service) VerificationEmail(ctx context.Context, r *pb.Authentication) (*pbn.Notification, error) {
	if err := validateUsername(r.GetUsername()); nil != err {
		return nil, err
	}

	rdb := util.NewRedis(1)
	otp := strconv.Itoa(util.Random(100000, 999999))
	keyRedis := fmt.Sprintf("otp_code_%s", r.GetUsername())
	otpRedis, err := util.GetRedis(rdb, keyRedis)
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("set otp to redis failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.VerifyEmailFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	if otpRedis != nil {
		otp = otpRedis.(string)
	}
	keyRedis = fmt.Sprintf("otp_code_%s", r.GetUsername())
	if err := util.SetRedis(rdb, keyRedis, otp, 10*time.Minute); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("set otp to redis failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.VerifyEmailFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	sender := util.NewGmailSender()
	err = sender.SendEmail("Verification Email", otp, []string{r.GetUsername()}, []string{}, []string{}, []string{})
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("set otp to redis failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.VerifyEmailFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	return &pbn.Notification{
		Status:  "OK",
		Message: constant.VerifyEmailSuccessfully,
		Code:    http.StatusOK,
	}, nil
}

func (s *Service) CheckingExpirationTime(ctx context.Context, r *pb.CheckingOtpExpirationRequest) (*pbn.Notification, error) {
	if err := validateCheckingOtpExpiration(r); nil != err {
		return nil, err
	}

	username := r.GetUsername()
	otp := r.GetOtp()

	rdb := util.NewRedis(1)
	keyRedis := "otp_code_" + username
	val, err := util.GetRedis(rdb, keyRedis)
	if err != nil && !errors.Is(err, redis.Nil) {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("get otp to redis failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.CheckingOtpExpirationFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	if val == nil {
		return &pbn.Notification{
			Status:  "Bad Request",
			Message: constant.ExpiredOtp,
			Code:    http.StatusBadRequest,
		}, nil
	}

	if val.(string) != otp {
		return &pbn.Notification{
			Status:  "Bad Request",
			Message: constant.InvalidOtp,
			Code:    http.StatusBadRequest,
		}, nil
	}

	if err := util.DeleteRedis(rdb, "otp_code"+r.GetUsername()); err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("delete otp to redis failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.CheckingOtpExpirationFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	return &pbn.Notification{
		Status:  "OK",
		Message: constant.CheckingOtpExpirationSuccessfully,
		Code:    http.StatusOK,
	}, nil
}

func (s *Service) CheckExsitUsername(ctx context.Context, r *pb.Authentication) (*pbn.Notification, error) {
	if err := validateUsername(r.GetUsername()); nil != err {
		return &pbn.Notification{
			Status:  "Bad Request",
			Message: constant.CheckingExistedUsernameFailed,
			Code:    http.StatusBadRequest,
		}, nil
	}

	count, err := NewRepository(s.db).checkExistUsername(r.GetUsername())
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "Authentication Service",
			"method":  "POST",
		}).WithError(err).Error("check existed username failed")
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: constant.CheckingExistedUsernameFailed,
			Code:    http.StatusInternalServerError,
		}, nil
	}

	if count > 0 {
		return &pbn.Notification{
			Status:  "Bad Request",
			Message: constant.ExistedUsername,
			Code:    http.StatusBadRequest,
		}, nil
	}

	return &pbn.Notification{
		Status:  "OK",
		Message: constant.CheckingExistedUsernameSuccessfully,
		Code:    http.StatusOK,
	}, nil
}

func (s *Service) CheckSessionExpired(ctx context.Context, r *types.Empty) (*pbn.Notification, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pbn.Notification{
			Status:  constant.CheckingSessionExpiredSuccessFailed,
			Message: "context null",
			Code:    http.StatusBadRequest,
		}, nil
	}

	jwtToken := md.Get("Authorization")
	if jwtToken[0] == "" {
		return &pbn.Notification{
			Status:  constant.CheckingSessionExpiredSuccessFailed,
			Message: "jwt_code is required!",
			Code:    http.StatusBadRequest,
		}, nil
	}

	token, err := jwt.JWTAuthService().ValidateToken(jwtToken[0])
	if err != nil || !token.Valid {
		return &pbn.Notification{
			Status:  constant.CheckingSessionExpiredSuccessFailed,
			Message: "jwt_code is invalid!",
			Code:    http.StatusBadRequest,
		}, nil
	}

	claims := token.Claims.(j.MapClaims)
	username, _ := claims["name"].(string)
	if username == "" {
		return &pbn.Notification{
			Status:  constant.CheckingSessionExpiredSuccessFailed,
			Message: "username is invalid",
			Code:    http.StatusInternalServerError,
		}, nil
	}

	rdb := util.NewRedis(2)
	keyRedis := username + "_jwt_token"
	_, err = util.GetRedis(rdb, keyRedis)
	if err != nil {
		return &pbn.Notification{
			Status:  constant.CheckingSessionExpiredSuccessFailed,
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}, nil
	}

	return &pbn.Notification{
		Status:  constant.CheckingSessionExpiredSuccessfully,
		Message: "OK",
		Code:    http.StatusOK,
	}, nil
}
