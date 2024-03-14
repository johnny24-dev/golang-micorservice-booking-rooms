package model

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name    string
	Price   float64
	HotelID int32
}
