package booked_room

import (
	model "github.com/nekizz/final-project/backend/booking/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) GetListBookedRoomOfBooking(bookingID int) ([]*model.BookedRoom, error) {
	var listBooked []*model.BookedRoom

	query := r.db.Model(&model.BookedRoom{}).Where("booking_id = ?", bookingID).Find(&listBooked)
	if err := query.Error; nil != err {
		return nil, err
	}

	return listBooked, nil
}
