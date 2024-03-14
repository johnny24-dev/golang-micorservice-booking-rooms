package util

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	modelB "github.com/nekizz/final-project/backend/booking/model"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"gorm.io/gorm"
)

func DecodeCreatePurchaseCmd(payload message.Payload) (*modelB.Booking, *pb.Booking, error) {
	var cmd pb.CreateBookingCmd
	if err := json.Unmarshal(payload, &cmd); err != nil {
		return nil, nil, err
	}

	bookingId := cmd.BookingId
	//
	//pbBookedRoom:= cmd.Booking.Order.PurchasedItems
	//var purchasedItems []model.PurchasedItem
	//for _, pbPurchasedItem := range pbPurchasedItems {
	//	purchasedItems = append(purchasedItems, model.PurchasedItem{
	//		ProductID: pbPurchasedItem.ProductId,
	//		Amount:    pbPurchasedItem.Amount,
	//	})
	//}
	return &modelB.Booking{
		Model: gorm.Model{
			ID: uint(bookingId),
		},
		BookDate:   nil,
		Note:       "",
		PaymentID:  0,
		CustomerID: 0,
	}, cmd.Booking, nil
}

//ID: purchaseID,
//Order: &model.Order{
//ID:             purchaseID,
//CustomerID:     cmd.Purchase.Order.CustomerId,
//PurchasedItems: &purchasedItems,
//},
//Payment: &model.Payment{
//ID:           purchaseID,
//CustomerID:   cmd.Purchase.Order.CustomerId,
//CurrencyCode: cmd.Purchase.Payment.CurrencyCode,
//Amount:       cmd.Purchase.Payment.Amount,
//},
