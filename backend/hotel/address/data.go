package address

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/address"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
	"time"
)

func PrepareAddressToRequest(p *pb.Address) *model.Address {
	return &model.Address{
		Model:         gorm.Model{},
		Province:      p.Province,
		DetailAddress: p.DetailAddress,
		Type:          p.Type,
	}
}

func PrepareDataToResponse(p *model.Address) *pb.Address {
	return &pb.Address{
		Id:            int32(p.ID),
		Province:      p.Province,
		DetailAddress: p.DetailAddress,
		Type:          p.Type,
	}
}

func getModel(id uint, c *model.Address) *model.Address {
	return &model.Address{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Province:      c.Province,
		DetailAddress: c.DetailAddress,
		Type:          c.Type,
	}
}
