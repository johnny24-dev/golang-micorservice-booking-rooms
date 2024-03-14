package payment

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/payment"
	"github.com/nekizz/final-project/backend/payment/model"
	"gorm.io/gorm"
)

func preparePaymentToRequest(p *pb.Payment) *model.Payment {
	return &model.Payment{
		Model:       gorm.Model{},
		PayemntType: p.PaymentType,
		CardName:    p.CardName,
		CardType:    p.CardType,
		CardNumber:  p.CardNumber,
		ExpiredDate: p.ExpiredDate,
	}
}

func preparePaymentToResponse(p *model.Payment) *pb.Payment {
	return &pb.Payment{
		Id:          int32(p.ID),
		PaymentType: "",
		CardName:    "",
		CardType:    "",
		CardNumber:  "",
		ExpiredDate: "",
		BookingId:   0,
	}
}
