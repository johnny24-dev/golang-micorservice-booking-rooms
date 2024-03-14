package payment

import (
	"github.com/nekizz/final-project/backend/payment/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatOne(e *model.Payment) (*model.Payment, error) {
	query := r.db.Model(&model.Payment{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}
