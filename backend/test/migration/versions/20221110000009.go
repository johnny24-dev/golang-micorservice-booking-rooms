package versions

import (
	"gorm.io/gorm"
)

func Version20221110000009(tx *gorm.DB) error {
	type Hotel struct {
		gorm.Model
	}

	type Customer struct {
		gorm.Model
		Hotel []*Hotel `gorm:"many2many:wishlists;"`
	}

	return tx.AutoMigrate(&Hotel{}, &Customer{})
}
