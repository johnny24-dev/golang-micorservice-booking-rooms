package account

import (
	"errors"
	"github.com/nekizz/final-project/backend/user/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Login(username string, password string) (*model.Account, error) {
	var account model.Account
	var count int64

	queryC := r.db.Model(&model.Account{}).Where("username = ? AND password = ?", username, password).Count(&count)
	if err := queryC.Error; nil != err {
		return nil, err
	}
	if count <= 0 {
		return nil, errors.New("Wrong username or password!")
	}

	query := r.db.Model(&model.Account{}).Where("username = ? AND password = ?", username, password).Find(&account)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &account, nil
}
