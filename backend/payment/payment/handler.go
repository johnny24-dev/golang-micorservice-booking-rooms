package payment

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

//func (s *Service) CreatePayment(msg *message.Message) ([]*message.Message, error) {
//	var replyMsgs []*message.Message
//
//	booking, pbBooking, err := util.DecodeCreatePurchaseCmd(msg.Payload)
//	if err != nil {
//		return nil, err
//	}
//	reply := pbBK.CreateBookingResponse{
//		BookingId: uint32(booking.ID),
//		Booking:   pbBooking,
//	}
//	_, err = NewRepository(s.db).CreatOne(preparePaymentToRequest(pbBooking.Payment))
//	if err != nil {
//		reply.Success = false
//		reply.Error = err.Error()
//	} else {
//		reply.Success = true
//		reply.Error = ""
//	}
//	reply.Timestamp = time.Now().String()
//
//	payload, err := json.Marshal(&reply)
//	if err != nil {
//		return nil, err
//	}
//	replyMsg := message.NewMessage(watermill.NewUUID(), payload)
//	replyMsgs = append(replyMsgs, replyMsg)
//
//	return replyMsgs, nil
//}

func RollbackPayment(msg *message.Message) error {

	return nil
}
