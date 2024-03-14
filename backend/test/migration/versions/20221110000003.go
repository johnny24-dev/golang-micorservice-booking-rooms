package versions

import (
	"gorm.io/gorm"
)

func Version20221110000003(tx *gorm.DB) error {
	type Address struct {
		gorm.Model
	}

	type Hotel struct {
		gorm.Model
		Name        string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		StarLevel   uint   `gorm:"TYPE:INT"`
		IsAvailable bool   `gorm:"TYPE:BOOLEAN"`
		Description string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		AddressID   uint
		CustomerID  uint
		Address     Address `gorm:"foreignKey:AddressID"`
	}

	type Customer struct {
		gorm.Model
		Hotel []Hotel `gorm:"foreignKey:CustomerID"`
	}

	return tx.AutoMigrate(&Hotel{}, &Address{}, &Customer{})
}
