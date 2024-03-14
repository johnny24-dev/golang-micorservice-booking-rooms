package versions

import (
	"gorm.io/gorm"
	"time"
)

func Version20221110000001(tx *gorm.DB) error {
	type BookedRoom struct {
		gorm.Model
		CheckIn   *time.Time
		CheckOut  *time.Time
		Price     float64 `gorm:"TYPE:DOUBLE"`
		Discount  float64 `gorm:"TYPE:DOUBLE"`
		IsCheckIn string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Number    uint
		BookingID uint
		RoomID    uint
	}

	return tx.AutoMigrate(&BookedRoom{})
}
