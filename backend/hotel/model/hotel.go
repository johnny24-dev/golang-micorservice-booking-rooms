package model

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name        string
	StarLevel   int
	Rate        float32 `gorm:"-"`
	IsAvailable bool
	Description string
	AddressID   int
	CustomerID  int
	IsDeleted   int
	Address     Address
	Image       []Image
	Comment     []Comment
}
