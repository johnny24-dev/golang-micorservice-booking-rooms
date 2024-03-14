package customer

import (
	"context"
	"gorm.io/gorm"
	"strconv"

	"github.com/gogo/protobuf/types"
	pb "github.com/nekizz/final-project/backend/go-pbtype/customer"
	"github.com/nekizz/final-project/backend/user/logger"
	userH "github.com/nekizz/final-project/backend/user/user"
	"github.com/sirupsen/logrus"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Customer) (*pb.Customer, error) {
	if err := validateCreateCustomer(r); nil != err {
		return nil, err
	}

	user, err := userH.CreateUserHandler(r, s.db)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("create customer failed")
		return nil, err
	}

	customerModel := prepareDataToRequest(r)
	customerModel.UserID = int(user.Id)
	customer, err := NewRepository(s.db).CreatOne(customerModel)
	if nil != err {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("create customer failed")
		return nil, err
	}

	customerResp := prepareDataToResponse(customer)
	customerResp.User = user

	return customerResp, nil
}

func (s Service) Update(ctx context.Context, r *pb.Customer) (*pb.Customer, error) {
	if err := validateUpdate(r); nil != err {
		return nil, err
	}

	customer, err := NewRepository(s.db).UpdateOne(int(r.GetId()), prepareDataToRequest(r))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "PUT",
		}).WithError(err).Error("update customer failed")
		return nil, err
	}

	user, err := userH.UpdateUserHandler(r, s.db)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "PUT",
		}).WithError(err).Error("update customer failed")
		return nil, err
	}

	customerPb := prepareDataToResponse(customer)
	customerPb.User = user

	return customerPb, nil
}

func (s Service) Get(ctx context.Context, r *pb.OneCustomerRequest) (*pb.Customer, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	id, _ := strconv.Atoi(r.GetId())
	customer, err := NewRepository(s.db).FindOne(id)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("get customer failed")
		return nil, err
	}

	user, err := userH.GetUserHandler(customer.UserID, s.db)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "GET",
		}).WithError(err).Error("get customer failed")
		return nil, err
	}

	customerPb := prepareDataToResponse(customer)
	customerPb.User = user

	return customerPb, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneCustomerRequest) (*types.Empty, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	id, _ := strconv.Atoi(r.GetId())
	err := NewRepository(s.db).DeleteOne(id)
	if nil != err {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("delete customer failed")
		return nil, err
	}

	return &types.Empty{}, nil
}
