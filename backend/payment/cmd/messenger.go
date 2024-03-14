package cmd

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/nekizz/final-project/backend/payment/constant"

	"gorm.io/gorm"
	"time"
)

func initMessenger(dbOrm *gorm.DB) (*kafka.Publisher, *message.Router, error) {
	brokers := []string{"103.184.113.181:9092"}
	watermillLogger := watermill.NewStdLogger(false, false)

	// Create Publisher
	publisherConfig := kafka.PublisherConfig{Brokers: brokers, Marshaler: kafka.DefaultMarshaler{}}
	publisher, err := kafka.NewPublisher(publisherConfig, watermillLogger)
	if err != nil {
		return nil, nil, err
	}

	// Declare Sarama Subscriber Config
	subscriberCfg := kafka.DefaultSaramaSubscriberConfig()
	subscriberCfg.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriberConfig := kafka.SubscriberConfig{
		Brokers:               brokers,
		Unmarshaler:           kafka.DefaultMarshaler{},
		OverwriteSaramaConfig: subscriberCfg,
		ConsumerGroup:         constant.PaymentGroup,
	}
	_, err = kafka.NewSubscriber(subscriberConfig, watermillLogger)
	if err != nil {
		return nil, nil, err
	}

	// Create Router
	router, err := message.NewRouter(message.RouterConfig{}, watermillLogger)
	if err != nil {
		return nil, nil, err
	}

	// Add Middleware and Plugin
	middlewareRetry := middleware.Retry{MaxRetries: 2, InitialInterval: time.Second * 10, Logger: watermillLogger}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(middleware.CorrelationID, middlewareRetry.Middleware, middleware.Recoverer)

	//paymentService := payment.NewService(dbOrm, publisher)
	//router.AddNoPublisherHandler("CreatePaymentHandler", constant.CreatePaymentTopic, subscriber, paymentService.CreatePayment)
	//router.AddNoPublisherHandler("RollbackPaymentHandler", constant.RollbackPaymentTopic, subscriber, payment.RollbackPayment)

	return publisher, router, err
}
