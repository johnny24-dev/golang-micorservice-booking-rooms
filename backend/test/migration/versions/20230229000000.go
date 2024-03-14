package versions

import "gorm.io/gorm"

func Version20230229000000(tx *gorm.DB) error {
	type Amenity struct {
		gorm.Model
		Name   string `gorm:"TYPE:VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci"`
		RoomID uint
	}

	type Room struct {
		gorm.Model
		Amenity         []Amenity `gorm:"foreignKey:RoomID"`
		Quantity        uint
		CurrentQuantity uint
	}

	return tx.AutoMigrate(&Amenity{}, &Room{})
}
