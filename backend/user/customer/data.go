package customer

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/customer"
	"github.com/nekizz/final-project/backend/user/model"
	"gorm.io/gorm"
)

func prepareDataToResponse(p *model.Customer) *pb.Customer {
	return &pb.Customer{
		Id:          int32(p.ID),
		Role:        p.Role,
		Description: p.Description,
	}
}

func prepareDataToRequest(p *pb.Customer) *model.Customer {
	return &model.Customer{
		Model:       gorm.Model{},
		Role:        p.Role,
		Description: p.Description,
	}
}

func getModel(id uint, c *model.Customer) *model.Customer {
	return &model.Customer{
		Model: gorm.Model{
			ID: id,
		},
		Role:        c.Role,
		Description: c.Description,
	}
}
