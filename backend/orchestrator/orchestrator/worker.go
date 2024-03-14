package orchestrator

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	modelBK "github.com/nekizz/final-project/backend/booking/model"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"github.com/nekizz/final-project/backend/orchestrator/constant"
	"github.com/nekizz/final-project/backend/orchestrator/logger"
	"github.com/nekizz/final-project/backend/orchestrator/model"
	"github.com/nekizz/final-project/backend/orchestrator/util"
	"time"
)

//rollback result
func (s *Service) publishBookingResult(ctx context.Context, booking *util.BookingResult, correlationID string) {
	result := util.EncodeDomainPurchaseResult(booking)
	payload, err := json.Marshal(result)
	if err != nil {
		logger.Log.WithError(err)
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	middleware.SetCorrelationID(correlationID, msg)
	if err := util.PublishMessage(s.publisher, constant.BookingResultTopic, msg, util.RESULT_MSG); err != nil {
		logger.Log.WithError(err)
	}
}

func (s *Service) publishRollbackResult(ctx context.Context, step string, rollbackResponse *model.RollbackResponse, correlationID string) {
	if !rollbackResponse.Success {
		s.publishBookingResult(ctx, &util.BookingResult{
			CustomerID: uint32(rollbackResponse.CustomerID),
			BookingID:  uint32(rollbackResponse.BookingID),
			Step:       step,
			Status:     util.StatusRollbackFailed,
		}, correlationID)
	}
}

//Room service
func (s *Service) rollbackUpdateRoomQuantity(ctx context.Context, customerID, bookingID uint64, correlationID string) error {
	s.publishBookingResult(ctx, &util.BookingResult{
		CustomerID: uint32(customerID),
		BookingID:  uint32(bookingID),
		Step:       util.StepUpdateRoomQuantity,
		Status:     util.StatusFailed,
		Timestamp:  time.Time{},
	}, correlationID)

	cmd := &pb.RollbackCmd{
		BookingId: bookingID,
		Timestamp: time.Now().String(),
	}
	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	middleware.SetCorrelationID(correlationID, msg)

	s.publishBookingResult(ctx, &util.BookingResult{}, correlationID)

	return util.PublishMessage(s.publisher, constant.RollbackRoomQuantityHandler, msg, util.TX_MSG)
}

//BookedRoom service
//func (s *Service) createBookedRoom(ctx context.Context, booking *pb.Booking, correlationID string) error {
//	s.publishBookingResult(ctx, &util.BookingResult{
//		CustomerID: uint32(booking.CustomerID),
//		BookingID:  uint32(booking.ID),
//		Step:       util.StepUpdateRoomQuantity,
//		Status:     util.StatusSucess,
//	}, correlationID)
//
//	cmd := util.EncodeDomainBooking(booking)
//	payload, err := json.Marshal(cmd)
//	if err != nil {
//		return err
//	}
//	msg := message.NewMessage(watermill.NewUUID(), payload)
//	middleware.SetCorrelationID(correlationID, msg)
//
//	s.publishBookingResult(ctx, &util.BookingResult{
//		CustomerID: uint32(booking.CustomerID),
//		BookingID:  uint32(booking.ID),
//		Step:       util.StepCreateBookedRoom,
//		Status:     util.StatusExecute,
//	}, correlationID)
//
//	return util.PublishMessage(s.publisher, constant.CreateBookedRoomHandler, msg, util.TX_MSG)
//}

func (s *Service) rollbackBookedRoom(bookingID uint64, correlationID string) error {
	cmd := &pb.RollbackCmd{
		BookingId: bookingID,
		Timestamp: time.Now().String(),
	}
	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	middleware.SetCorrelationID(correlationID, msg)
	return util.PublishMessage(s.publisher, constant.RollbackBookedRoomHandler, msg, util.TX_MSG)
}

//Payment service
func (s *Service) createPayment(ctx context.Context, booking *modelBK.Booking, correlationID string) error {
	s.publishBookingResult(ctx, &util.BookingResult{
		CustomerID: uint32(booking.CustomerID),
		BookingID:  uint32(booking.ID),
		Step:       util.StepUpdateRoomQuantity,
		Status:     util.StatusSucess,
	}, correlationID)

	cmd := util.EncodeDomainBooking(booking)
	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	middleware.SetCorrelationID(correlationID, msg)

	s.publishBookingResult(ctx, &util.BookingResult{
		CustomerID: uint32(booking.CustomerID),
		BookingID:  uint32(booking.ID),
		Step:       util.StepUpdateRoomQuantity,
		Status:     util.StatusSucess,
	}, correlationID)

	return util.PublishMessage(s.publisher, constant.CreatePaymentHandler, msg, util.TX_MSG)
}

func (s *Service) rollbackPayment(bookingID uint64, correlationID string) error {
	cmd := &pb.RollbackCmd{
		BookingId: bookingID,
		Timestamp: time.Now().String(),
	}
	payload, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	msg := message.NewMessage(watermill.NewUUID(), payload)
	middleware.SetCorrelationID(correlationID, msg)
	return util.PublishMessage(s.publisher, constant.RollbackPaymentHandler, msg, util.TX_MSG)
}
