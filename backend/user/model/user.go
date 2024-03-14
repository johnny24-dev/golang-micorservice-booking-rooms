package model

import (
	address "github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Avatar    string
	Name      string
	Email     string
	Phone     string
	About     string
	Note      string
	Gender    string
	AddressID int32
	AccountID int32
	Address   address.Address
}
