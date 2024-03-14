package authentication

import (
	"gorm.io/gorm"

	"github.com/nekizz/final-project/backend/authentication/model"

	pba "github.com/nekizz/final-project/backend/go-pbtype/authentication"
)

func prepareDataToResponse(p *model.Account) *pba.Authentication {
	return &pba.Authentication{
		Id:       int32(p.ID),
		Username: p.Username,
	}
}

func prepareDataToRequest(p *pba.Authentication) *model.Account {
	return &model.Account{
		Model:    gorm.Model{},
		Username: p.Username,
		Password: p.Password,
	}
}

func getModel(id uint, c *model.Account) *model.Account {
	return &model.Account{
		Model: gorm.Model{
			ID: id,
		},
		Username: c.Username,
		Password: c.Password,
	}
}
