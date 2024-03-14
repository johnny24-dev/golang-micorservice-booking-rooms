package modelDecodeCreatePurchaseCmd

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	Status     string
	Note       string
	StartDate  *time.Time
	EndDate    *time.Time
	CustomerID uint
}

type BookingV2 struct {
	gorm.Model
	Status     string
	Note       string
	StartDate  *time.Time
	EndDate    *time.Time
	CustomerID uint
}
