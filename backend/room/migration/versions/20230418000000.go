package versions

import "gorm.io/gorm"

func Version20230418000000(tx *gorm.DB) error {
	type Room struct {
		gorm.Model
		CurrentQuantity uint
	}

	return tx.AutoMigrate(&Room{})
}
