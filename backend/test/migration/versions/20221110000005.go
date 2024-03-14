package versions

import (
	"gorm.io/gorm"
)

func Version20221110000005(tx *gorm.DB) error {
	type Room struct {
		gorm.Model
		Name        string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Type        string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Price       float64 `gorm:"TYPE:FLOAT"`
		Status      string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Description string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		HotelID     uint
	}

	type Hotel struct {
		gorm.Model
		Room []Room `gorm:"foreignKey:HotelID"`
	}

	return tx.AutoMigrate(&Room{}, &Hotel{})
}
