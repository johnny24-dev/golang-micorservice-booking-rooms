package versions

import (
	"gorm.io/gorm"
)

func Version20221110000006(tx *gorm.DB) error {
	type Service struct {
		gorm.Model
		Name    string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Price   float64 `gorm:"TYPE:DOUBLE"`
		HotelID uint
	}

	type Hotel struct {
		gorm.Model
		Service []Service `gorm:"foreignKey:HotelID"`
	}

	return tx.AutoMigrate(&Service{}, &Hotel{})
}
