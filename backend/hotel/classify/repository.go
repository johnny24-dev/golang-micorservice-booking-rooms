package classify

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

func (r *Repository) getAllClassify(listId []int32) ([]*model.Classify, error) {
	var list []*model.Classify

	query := r.db.Model(&model.Classify{}).Where("id IN ?", listId).Find(&list)
	if err := query.Error; nil != err {
		return nil, err
	}

	return list, nil
}
