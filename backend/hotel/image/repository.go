package image

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

func (r *Repository) ListAllImage(hotelID int) ([]*model.Image, error) {
	var list []*model.Image

	query := r.db.Model(&model.Image{}).Select("*").Where("hotel_id = ?", hotelID)
	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Repository) UpdateImage(image *model.Image, id int) (*model.Image, error) {
	query := r.db.Model(&model.Image{}).Select("*").Omit("hotel_id").Where("id = ?", id).UpdateColumns(&image)
	if err := query.Error; err != nil {
		return nil, err
	}

	return image, nil
}
