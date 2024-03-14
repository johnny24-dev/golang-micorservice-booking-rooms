package user

import (
	"fmt"
	"github.com/nekizz/final-project/backend/user/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpdateOne(id int, c *model.User) (*model.User, error) {
	query := r.db.Model(&model.User{}).Where("id = ?", id).UpdateColumns(getModel(uint(id), c))
	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}

func (r *Repository) UpdateAddressID(idUser int, c *model.User) (*model.User, error) {
	fmt.Println("address_id ", c.Address.ID)
	query := r.db.Model(&model.User{}).Where("id = ?", idUser).Updates(map[string]interface{}{"address_id": c.Address.ID})
	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}

func (r *Repository) FindOne(id int) (*model.User, error) {
	var user model.User

	query := r.db.Model(&model.User{}).Where("id = ?", id).Find(&user)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) CreateOne(c *model.User) (*model.User, error) {
	query := r.db.Model(&model.User{}).Create(c)
	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}

func (r *Repository) DeleteOne(id int) error {
	query := r.db.Delete(&model.User{}, "id = ?", id)
	if err := query.Error; nil != err {
		return err
	}

	return nil
}
