package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Province      string
	DetailAddress string
	Type          string
}
