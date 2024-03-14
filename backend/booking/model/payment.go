package modelDecodeCreatePurchaseCmd

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	PaypalPaymentID string
	Status          string
	BookingID       uint
}
