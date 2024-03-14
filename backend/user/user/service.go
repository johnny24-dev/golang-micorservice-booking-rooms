package user

import (
	"context"
	"gorm.io/gorm"

	"github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"

	pb "github.com/nekizz/final-project/backend/go-pbtype/user"
	"github.com/nekizz/final-project/backend/user/logger"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Update(ctx context.Context, r *pb.User) (*pb.User, error) {
	if err := validateUpdate(r); err != nil {
		return nil, err
	}

	user, err := NewRepository(s.db).UpdateOne(int(r.GetId()), prepareUserToRequest(r))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "PUT",
		}).WithError(err).Error("update user failed")
		return nil, err
	}

	return prepareUserToResponse(user), nil
}

func (s *Service) Delete(ctx context.Context, r *pb.OneUserRequest) (*types.Empty, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	err := NewRepository(s.db).DeleteOne(int(r.GetId()))
	if nil != err {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "DEL",
		}).WithError(err).Error("delete user failed")
		return nil, err
	}

	return &types.Empty{}, nil
}
