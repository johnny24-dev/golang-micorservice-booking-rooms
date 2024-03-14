package versions

import "gorm.io/gorm"

func Version20230101000000(tx *gorm.DB) error {
	type Image struct {
		gorm.Model
		Number  uint   `gorm:"TYPE:INT"`
		Url     string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Type    string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		HotelID uint
	}

	type Hotel struct {
		gorm.Model
		Image []Image `gorm:"foreignKey:HotelID"`
	}

	return tx.AutoMigrate(&Hotel{}, &Image{})
}
