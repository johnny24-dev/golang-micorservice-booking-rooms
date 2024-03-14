package versions

import (
	"gorm.io/gorm"
)

func Version20230529000000(tx *gorm.DB) error {
	type Permission struct {
		gorm.Model
		HotelID          uint
		CustomerID       uint
		AccessPermission uint
	}

	return tx.AutoMigrate(&Permission{})
}
