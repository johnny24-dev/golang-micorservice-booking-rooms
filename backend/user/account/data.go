package account

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/account"
	"github.com/nekizz/final-project/backend/user/model"
	"gorm.io/gorm"
)

func PrepareDataToResponse(p *model.Account) *pb.Account {
	return &pb.Account{
		Id:       int32(p.ID),
		Username: p.Username,
		Password: p.Password,
	}
}

func PrepareDataToRequest(p *pb.Account) *model.Account {
	return &model.Account{
		Model:    gorm.Model{},
		Username: p.Username,
		Password: p.Password,
	}
}

func getModel(id uint, c *model.Account) *model.Account {
	c.ID = id
	return &model.Account{
		Model:    gorm.Model{},
		Username: c.Username,
		Password: c.Password,
	}
}
