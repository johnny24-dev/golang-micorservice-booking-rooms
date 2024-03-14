package versions

import (
	"gorm.io/gorm"
	"time"
)

func Version20221110000002(tx *gorm.DB) error {
	type Booking struct {
		gorm.Model
		BookDate   *time.Time
		Total      float64 `gorm:"TYPE:DOUBLE"`
		Status     string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		Note       string  `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		CustomerID uint
	}

	return tx.AutoMigrate(&Booking{})
}
