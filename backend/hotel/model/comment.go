package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text       string
	Type       string
	Rate       float32
	HotelID    int
	CustomerID int
}
