package wishlist

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

func (r *Repository) CreatOne(e *model.Wishlist) (*model.Wishlist, error) {
	query := r.db.Model(&model.Wishlist{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) DeleteOne(hotelID int, customerID int) error {
	query := r.db.Delete(&model.Wishlist{}, "hotel_id = ? AND customer_id = ?", hotelID, customerID)
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func (r *Repository) ListWishlistIdOfCustomer(customerID int) ([]int, int64, error) {
	var count int64
	var listWishlistID []int

	listQuery := r.db.Model(&model.Wishlist{}).Select("hotel_id").Where("customer_id = ?", customerID).Find(&listWishlistID)
	if err := listQuery.Error; nil != err {
		return nil, 0, err
	}
	countQuery := r.db.Model(&model.Wishlist{}).Select("hotel_id").Where("customer_id = ?", customerID)
	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	return listWishlistID, count, nil
}
