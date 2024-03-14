package cmd

//func initMessenger(dbOrm *gorm.DB) (*kafka.Publisher, *message.Router, error) {
//	brokers := []string{"103.184.113.181:9092"}
//	watermillLogger := watermill.NewStdLogger(false, false)
//
//	// Create Publisher
//	publisherConfig := kafka.PublisherConfig{Brokers: brokers, Marshaler: kafka.DefaultMarshaler{}}
//	publisher, err := kafka.NewPublisher(publisherConfig, watermillLogger)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	// Declare Sarama Subscriber Config
//	subscriberCfg := kafka.DefaultSaramaSubscriberConfig()
//	subscriberCfg.Consumer.Offsets.Initial = sarama.OffsetOldest
//
//	subscriberConfig := kafka.SubscriberConfig{
//		Brokers:               brokers,
//		Unmarshaler:           kafka.DefaultMarshaler{},
//		OverwriteSaramaConfig: subscriberCfg,
//		ConsumerGroup:         constant.RoomGroup,
//	}
//	subscriber, err := kafka.NewSubscriber(subscriberConfig, watermillLogger)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	// Create Router
//	router, err := message.NewRouter(message.RouterConfig{}, watermillLogger)
//	if err != nil {
//		return nil, nil, err
//	}
//
//	// Add Middleware and Plugin
//	middlewareRetry := middleware.Retry{MaxRetries: 2, InitialInterval: time.Second * 10, Logger: watermillLogger}
//	router.AddPlugin(plugin.SignalsHandler)
//	router.AddMiddleware(middleware.CorrelationID, middlewareRetry.Middleware, middleware.Recoverer)
//
//	roomService := room.NewService(dbOrm, publisher)
//	//router.AddHandler("sagaroom_update_quantity_room_handler", constant.UpdateRoomQuantityHandler, subscriber, constant.ReplyTopic, publisher, roomService.UpdateRoomQuantity)
//	//router.AddHandler("sagaroom_rollback_quantity_room_handler", constant.RollbackRoomQuantityTopic, subscriber, constant.ReplyTopic, publisher, roomService.RollbackRoomQuantity)
//
//	return publisher, router, err
//}
