package comment

import "gorm.io/gorm"

func CalculateRateOfHotel(hotelID int, db *gorm.DB) ([]float32, error) {
	listRate, err := NewRepository(db).ListRateOfHotel(hotelID)
	if err != nil {
		return nil, err
	}

	return listRate, nil
}
