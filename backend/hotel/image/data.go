package image

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/image"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
)

func PrepareImageToRequest(p *pb.Image) *model.Image {
	return &model.Image{
		Model:   gorm.Model{},
		HotelID: p.HotelId,
		Url:     p.Url,
		Type:    p.Type,
	}
}

func PrepareImageToResponse(p *model.Image) *pb.Image {
	return &pb.Image{
		Id:   int32(p.ID),
		Url:  p.Url,
		Type: p.Type,
	}
}

func getImageModel(id int, p *model.Image) *model.Image {
	return &model.Image{
		Model: gorm.Model{
			ID: uint(id),
		},
		Url:     p.Url,
		HotelID: p.HotelID,
		Type:    p.Type,
	}
}
