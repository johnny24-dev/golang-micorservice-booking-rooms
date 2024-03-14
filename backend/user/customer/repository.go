package customer

import (
	"github.com/nekizz/final-project/backend/user/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatOne(e *model.Customer) (*model.Customer, error) {
	query := r.db.Model(&model.Customer{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) UpdateOne(id int, c *model.Customer) (*model.Customer, error) {
	query := r.db.Model(&model.Customer{}).Where("id = ?", id).UpdateColumns(getModel(uint(id), c))
	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}

func (r *Repository) FindOne(id int) (*model.Customer, error) {
	var customer model.Customer

	query := r.db.Model(&model.Customer{}).Where("id = ?", id).Find(&customer)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &customer, nil
}

func (r *Repository) FindOneByAccountID(id int) (*model.Customer, error) {
	var customer model.Customer
	var userID int

	queryUser := r.db.Model(&model.User{}).Select("id").Where("account_id = ?", id).Find(&userID)
	if err := queryUser.Error; nil != err {
		return nil, err
	}

	query := r.db.Model(&model.Customer{}).Preload("User").Preload("User.Account").Preload("User.Address").Where("user_id = ?", userID).Find(&customer)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &customer, nil
}

func (r *Repository) DeleteOne(id int) error {
	query := r.db.Select("User").Delete(&model.Customer{}, "id = ?", id)
	if err := query.Error; nil != err {
		return err
	}

	return nil
}
