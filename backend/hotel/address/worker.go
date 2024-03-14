package address

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/address"
	"github.com/nekizz/final-project/backend/hotel/logger"
	"github.com/nekizz/final-project/backend/hotel/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func UpdateAddressesWorker(r *pb.Address, tx *gorm.DB) (*model.Address, error) {
	address := getModel(uint(r.GetId()), PrepareAddressToRequest(r))
	addresses, err := NewRepository(tx).UpdateOne(int(address.ID), address)
	if nil != err {
		logger.Log.WithFields(logrus.Fields{
			"service": "address",
			"method":  "CreateAddressesWorker",
			"data":    address,
		}).WithError(err).Error("Create customer addresses failed")

		return nil, err
	}

	return addresses, nil
}
