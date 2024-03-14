package versions

import "gorm.io/gorm"

func Version20230225000001(tx *gorm.DB) error {
	type Customer struct {
		gorm.Model
	}
	type Comment struct {
		gorm.Model
		Text       string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Type       string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Rate       float32
		HotelID    int
		CustomerID int
		Customer   Customer `gorm:"foreignKey:CustomerID"`
	}
	type Hotel struct {
		gorm.Model
		Comment []Comment `gorm:"foreignKey:HotelID"`
	}

	return tx.AutoMigrate(&Comment{}, &Customer{}, &Hotel{})
}
