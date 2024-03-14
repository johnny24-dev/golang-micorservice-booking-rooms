package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Role        string
	Description string
	UserID      int
}
