package user

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	pbe "github.com/nekizz/final-project/backend/go-pbtype/address"
	pb "github.com/nekizz/final-project/backend/go-pbtype/customer"
	pbu "github.com/nekizz/final-project/backend/go-pbtype/user"
	"github.com/nekizz/final-project/backend/user/logger"
)

func UpdateUserHandler(p *pb.Customer, db *gorm.DB) (*pbu.User, error) {
	var addressNew *pbe.Address
	userPb := p.GetUser()

	if userPb.Address != nil {
		addressAddr := viper.GetString("hotel_backend")
		addressConn, err := grpc.Dial(addressAddr, grpc.WithInsecure())
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"service": "User Service",
				"method":  "PUT",
			}).WithError(err).Error("dial hotel service failed")
			return nil, err
		}
		addressService := pbe.NewAddressServiceClient(addressConn)

		address := userPb.Address
		adrrReq := &pbe.Address{
			Id:            address.Id,
			Province:      address.Province,
			DetailAddress: address.DetailAddress,
			Type:          "customer",
		}

		if address.Id != 0 {
			addressNew, err = addressService.Update(context.Background(), adrrReq)
			if nil != err {
				logger.Log.WithFields(logrus.Fields{
					"service": "User Service",
					"method":  "PUT",
				}).WithError(err).Error("update address failed")
				return nil, err
			}
		} else {
			addressNew, err = addressService.Create(context.Background(), adrrReq)
			if nil != err {
				logger.Log.WithFields(logrus.Fields{
					"service": "User Service",
					"method":  "PUT",
				}).WithError(err).Error("create address failed")
				return nil, err
			}

			userPb.AddressId = addressNew.Id
		}
	}

	user, err := NewRepository(db).UpdateOne(int(userPb.GetId()), prepareUserToRequest(userPb))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "PUT",
		}).WithError(err).Error("update user handler failed")
		return nil, err
	}

	userPbNew := prepareUserToResponse(user)
	userPbNew.Address = addressNew

	return userPbNew, nil
}

func CreateUserHandler(p *pb.Customer, db *gorm.DB) (*pbu.User, error) {
	var addressNew *pbe.Address
	userPb := p.GetUser()

	if userPb.GetAddress() != nil {
		addressAddr := viper.GetString("hotel_backend")
		addressConn, err := grpc.Dial(addressAddr, grpc.WithInsecure())
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"service": "User Service",
				"method":  "POST",
			}).WithError(err).Error("dial address service failed")
			return nil, err
		}

		address := userPb.Address
		adrrReq := &pbe.Address{
			Id:            address.Id,
			Province:      address.Province,
			DetailAddress: address.DetailAddress,
			Type:          "customer",
		}
		addressService := pbe.NewAddressServiceClient(addressConn)
		addressNew, err = addressService.Create(context.Background(), adrrReq)
		if nil != err {
			logger.Log.WithFields(logrus.Fields{
				"service": "User Service",
				"method":  "POST",
			}).WithError(err).Error("create address failed")
			return nil, err
		}

		userPb.AddressId = addressNew.Id
	}

	user, err := NewRepository(db).CreateOne(prepareUserToRequest(userPb))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("create user handler failed")
		return nil, err
	}

	userPbNew := prepareUserToResponse(user)
	userPbNew.Address = addressNew

	return userPbNew, nil
}

func GetUserHandler(userID int, db *gorm.DB) (*pbu.User, error) {
	var addressNew *pbe.Address

	user, err := NewRepository(db).FindOne(userID)
	if err != nil {
		logger.Log.WithFields(logrus.Fields{
			"service": "User Service",
			"method":  "POST",
		}).WithError(err).Error("create user handler failed")
		return nil, err
	}

	if user.AddressID != 0 {
		addressAddr := viper.GetString("hotel_backend")
		addressConn, err := grpc.Dial(addressAddr, grpc.WithInsecure())
		if err != nil {
			logger.Log.WithFields(logrus.Fields{
				"service": "User Service",
				"method":  "PUT",
			}).WithError(err).Error("dial hotel service failed")
			return nil, err
		}
		addressService := pbe.NewAddressServiceClient(addressConn)
		adrrReq := &pbe.OneAddressRequest{
			Id: user.AddressID,
		}
		addressNew, err = addressService.Get(context.Background(), adrrReq)
		if nil != err {
			logger.Log.WithFields(logrus.Fields{
				"service": "User Service",
				"method":  "GET",
			}).WithError(err).Error("get address failed")
			return nil, err
		}
	}

	userPbNew := prepareUserToResponse(user)
	userPbNew.Address = addressNew

	return userPbNew, nil
}
