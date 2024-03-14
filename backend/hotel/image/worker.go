package image

import (
	"fmt"
	pb "github.com/nekizz/final-project/backend/go-pbtype/image"
	"github.com/nekizz/final-project/backend/hotel/logger"
	"github.com/nekizz/final-project/backend/hotel/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetImageWorker(hotelID int, tx *gorm.DB) ([]*model.Image, error) {
	listImage, err := NewRepository(tx).ListAllImage(hotelID)
	if nil != err {
		logger.Log.WithFields(logrus.Fields{
			"service": "image",
			"method":  "GetListImageWorker",
			"data":    listImage,
		}).WithError(err).Error("Get list image failed")
		return nil, err
	}
	fmt.Println(listImage)
	return listImage, nil
}

func UpdateImageWorker(r *pb.Image, tx *gorm.DB) (*model.Image, error) {
	data := getImageModel(int(r.GetId()), PrepareImageToRequest(r))
	image, err := NewRepository(tx).UpdateImage(data, int(r.GetId()))
	if nil != err {
		logger.Log.WithFields(logrus.Fields{
			"service": "image",
			"method":  "UpdateImageWorker",
			"data":    image,
		}).WithError(err).Error("Update image failed")
		return nil, err
	}

	return image, nil
}
