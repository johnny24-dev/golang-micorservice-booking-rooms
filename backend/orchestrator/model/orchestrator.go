package model

import (
	modelBK "github.com/nekizz/final-project/backend/booking/model"
	"time"
)

// CreatePurchaseResponse value object
type CreateBookingResponse struct {
	Booking *modelBK.Booking
	Success bool
	Error   string
}

// RollbackResponse value object
type RollbackResponse struct {
	CustomerID uint64
	BookingID  uint64
	Success    bool
	Error      string
}

type BookingResult struct {
	CustomerID uint64
	BookingID  uint64
	Step       string
	Status     string
	Timestamp  time.Time
}
