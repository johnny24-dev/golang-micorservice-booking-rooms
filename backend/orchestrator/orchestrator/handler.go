package orchestrator

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/nekizz/final-project/backend/orchestrator/util"
)

func (s *Service) StartBookingTransaction(msg *message.Message) error {
	fmt.Println("orchestrator service here")
	booking, err := util.DecodeCreateBookingCmd(msg.Payload)
	if err != nil {
		return err
	}
	correlationID := msg.Metadata.Get(middleware.CorrelationIDMetadataKey)
	return s.StartTransaction(context.Background(), booking, correlationID)
	return nil
}

func (s *Service) HandleBookingReply(msg *message.Message) error {
	correlationID := msg.Metadata.Get(middleware.CorrelationIDMetadataKey)
	return s.HandleReply(msg, correlationID)
	return nil
}
