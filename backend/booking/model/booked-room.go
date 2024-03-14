package modelDecodeCreatePurchaseCmd

import (
	"gorm.io/gorm"
	"time"
)

type BookedRoom struct {
	gorm.Model
	CheckIn   *time.Time
	CheckOut  *time.Time
	Price     float32
	IsCheckIn string
	Quantity  uint
	BookingID uint
	RoomID    uint
}
