package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"github.com/nekizz/final-project/backend/orchestrator/constant"
	"github.com/nekizz/final-project/backend/orchestrator/logger"
	"github.com/nekizz/final-project/backend/orchestrator/util"
	"gorm.io/gorm"
)

type Service struct {
	publisher *kafka.Publisher
	db        *gorm.DB
}

func NewService(db *gorm.DB, publisher *kafka.Publisher) *Service {
	return &Service{
		db:        db,
		publisher: publisher,
	}
}

func (s *Service) StartTransaction(parentCtx context.Context, booking *pb.Booking, correlationID string) error {
	fmt.Println("orchestrator service run")
	payload, err := json.Marshal(booking)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	middleware.SetCorrelationID(correlationID, msg)
	logger.Log.Info("create payment of booking: %v", booking.GetId())
	return util.PublishMessage(s.publisher, constant.PaymentTopic, msg, util.TX_MSG)
}

func (s *Service) HandleReply(msg *message.Message, correlationID string) error {
	handler := msg.Metadata.Get(constant.HandlerHeader)
	switch handler {
	//case constant.CreateBookedRoomHandler:
	//	resp, err := util.DecodeCreateBookingResponse(msg.Payload)
	//	if err != nil {
	//		return err
	//	}
	//	if resp.Success {
	//		s.publishBookingResult(ctx, &event.PurchaseResult{
	//			CustomerID: resp.Purchase.Order.CustomerID,
	//			PurchaseID: resp.Purchase.ID,
	//			Step:       event.StepCreatePayment,
	//			Status:     event.StatusSucess,
	//		}, correlationID)
	//
	//		return nil
	//	}
	//	return s.rollbackBookedRoom(uint64(resp.Booking.ID), correlationID)
	//case constant.RollbackBookedRoomHandler:
	//	resp, err := util.DecodeRollbackResponse(msg.Payload)
	//	if err != nil {
	//		return err
	//	}
	//	s.publishRollbackResult(context.Background(), util.StepCreateBookedRoom, resp, correlationID)
	case constant.CreatePaymentHandler:
		resp, err := util.DecodeCreateBookingResponse(msg.Payload)
		if err != nil {
			return err
		}
		if resp.Success {
			//return s.createBookedRoom(context.Background(), resp.Booking, correlationID)
			fmt.Println("come here")
		}
		return s.rollbackPayment(uint64(resp.Booking.GetId()), correlationID)
	case constant.RollbackPaymentHandler:
		resp, err := util.DecodeRollbackResponse(msg.Payload)
		if err != nil {
			return err
		}
		s.publishRollbackResult(context.Background(), util.StepCreatePayment, resp, correlationID)
	default:
		return fmt.Errorf("unkown handler: %s", handler)
	}
	return nil
}
