package util

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	modelBK "github.com/nekizz/final-project/backend/booking/model"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"github.com/nekizz/final-project/backend/orchestrator/model"
)

func DecodeCreateBookingCmd(payload message.Payload) (*pb.Booking, error) {
	var cmd pb.CreateBookingCmd

	if err := json.Unmarshal(payload, &cmd); err != nil {
		return nil, err
	}

	return cmd.Booking, nil
}

func EncodeDomainBooking(booking *modelBK.Booking) *pb.CreateBookingCmd {
	//var pbPurchasedItems []*pb.PurchasedItem
	//for _, purchasedItem := range *purchase.Order.PurchasedItems {
	//	pbPurchasedItems = append(pbPurchasedItems, &pb.PurchasedItem{
	//		ProductId: purchasedItem.ProductID,
	//		Amount:    purchasedItem.Amount,
	//	})
	//}
	cmd := &pb.CreateBookingCmd{
		//PurchaseId: purchase.ID,
		//Purchase: &pb.Purchase{
		//	Order: &pb.Order{
		//		CustomerId:     purchase.Order.CustomerID,
		//		PurchasedItems: pbPurchasedItems,
		//	},
		//	Payment: &pb.Payment{
		//		CurrencyCode: purchase.Payment.CurrencyCode,
		//		Amount:       purchase.Payment.Amount,
		//	},
		//},
		//Timestamp: pkg.Time2pbTimestamp(time.Now()),
	}
	return cmd
}

func EncodeDomainPurchaseResult(bookingResult *BookingResult) *pb.BookingResult {
	return &pb.BookingResult{
		//CustomerId: purchaseResult.CustomerID,
		//PurchaseId: purchaseResult.PurchaseID,
		//Step:       getPbPurchaseStep(purchaseResult.Step),
		//Status:     getPbPurchaseStatus(purchaseResult.Status),
		//Timestamp:  pkg.Time2pbTimestamp(time.Now()),
	}
}

func DecodeCreateBookingResponse(payload message.Payload) (*pb.CreateBookingResponse, error) {
	var resp *pb.CreateBookingResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func DecodeRollbackResponse(payload message.Payload) (*model.RollbackResponse, error) {
	var resp pb.RollbackResponse
	if err := json.Unmarshal(payload, &resp); err != nil {
		return nil, err
	}
	return &model.RollbackResponse{
		CustomerID: resp.CustomerId,
		BookingID:  resp.BookingId,
		Success:    resp.Success,
		Error:      resp.Error,
	}, nil
}
