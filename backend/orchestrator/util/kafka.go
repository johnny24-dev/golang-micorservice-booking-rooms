package util

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

type MessageType int

const (
	TX_MSG MessageType = iota
	RESULT_MSG
)

func PublishMessage(publisher *kafka.Publisher, topic string, msg *message.Message, messageType MessageType) error {
	switch messageType {
	case TX_MSG:
		return publisher.Publish(topic, msg)
	case RESULT_MSG:
		return publisher.Publish(topic, msg)
	default:
		return fmt.Errorf("unkown message type: %v", messageType)
	}
}
