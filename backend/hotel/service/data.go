package service

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/service"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
	"time"
)

func prepareDataToResponse(p *model.Service) *pb.Service {
	return &pb.Service{
		Id:    int32(p.ID),
		Name:  p.Name,
		Price: float32(p.Price),
	}
}

func prepareDataToRequest(p *pb.Service) *model.Service {
	return &model.Service{
		Model:   gorm.Model{},
		Name:    p.Name,
		Price:   float64(p.Price),
		HotelID: p.HotelId,
	}
}

func getModel(id uint, c *model.Service) *model.Service {
	return &model.Service{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:  c.Name,
		Price: c.Price,
	}
}
