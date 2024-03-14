package user

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/user"
	"github.com/nekizz/final-project/backend/user/model"
	"gorm.io/gorm"
	"time"
)

func prepareUserToResponse(p *model.User) *pb.User {
	return &pb.User{
		Id:        int32(p.ID),
		Name:      p.Name,
		Email:     p.Email,
		Phone:     p.Phone,
		Note:      p.Note,
		Avatar:    p.Avatar,
		Gender:    p.Gender,
		AccountId: p.AccountID,
	}
}

func prepareUserToRequest(p *pb.User) *model.User {
	return &model.User{
		Model:     gorm.Model{},
		Avatar:    p.Avatar,
		Name:      p.Name,
		Email:     p.Email,
		Phone:     p.Phone,
		Note:      p.Note,
		Gender:    p.Gender,
		AccountID: p.AccountId,
		AddressID: p.AddressId,
	}
}

func getModel(id uint, c *model.User) *model.User {
	c.ID = id
	return &model.User{
		Model: gorm.Model{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Avatar:  c.Avatar,
		Name:    c.Name,
		Email:   c.Email,
		Address: c.Address,
		Phone:   c.Phone,
		About:   c.About,
		Note:    c.Note,
	}
}
