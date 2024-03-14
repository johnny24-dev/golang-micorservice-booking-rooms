package versions

import "gorm.io/gorm"

func Version20230428000000(tx *gorm.DB) error {
	type Payment struct {
		gorm.Model
		PaypalPaymentID string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Status          string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		BookingID       uint
	}

	return tx.AutoMigrate(&Payment{})
}
