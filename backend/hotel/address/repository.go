package address

import (
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindOne(id int) (*model.Address, error) {
	var address model.Address

	query := r.db.Model(&model.Address{}).Where("id = ?", id).Find(&address)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &address, nil
}

func (r *Repository) CreatOne(e *model.Address) (*model.Address, error) {
	query := r.db.Model(&model.Address{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) UpdateOne(id int, c *model.Address) (*model.Address, error) {
	query := r.db.Model(&model.Address{}).Where("id = ?", id).UpdateColumns(getModel(uint(id), c))
	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}
