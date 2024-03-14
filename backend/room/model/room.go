package model

import "gorm.io/gorm"

type Room struct {
	gorm.Model
	Name            string
	Type            string
	Price           float32
	Status          string
	Description     string
	Quantity        uint
	CurrentQuantity uint
	Amenity         []Amenity
	HotelID         int32
}
