package constant

var (
	//Group
	OrchestratorGroup = "orchestrator_group"
	HandlerHeader     = "Handler"

	BookingTopic       = "booking.topic"
	BookingResultTopic = "booking.result"

	PaymentTopic = "payment.topic"

	UpdateRoomQuantityHandler   = "update_room_quantity_handler"
	RollbackRoomQuantityHandler = "rollback_room_quantity_handler"

	CreateBookedRoomHandler   = "create_booked_room_handler"
	RollbackBookedRoomHandler = "rollback_room_quantity_handler"

	CreatePaymentHandler   = "create_payment_handler"
	RollbackPaymentHandler = "rollback_payment_handler"

	UpdateRoomQuantityTopic = "update.room.quantity.topic"
	ReplyTopic              = "reply.topic"
)
