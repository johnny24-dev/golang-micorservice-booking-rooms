package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	PayemntType string
	CardName    string
	CardType    string
	CardNumber  string
	ExpiredDate string
	BookingID   uint
}
