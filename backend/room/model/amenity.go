package model

import "gorm.io/gorm"

type Amenity struct {
	gorm.Model
	Name   string
	RoomID uint
}
