package util

import "time"

var (
	StepUpdateRoomQuantity = "UPDATE_ROOM_QUANTITY"
	StepCreateBookedRoom   = "CREATE_BOOKED_ROOM"
	StepCreateOrder        = "CREATE_ORDER"
	StepCreatePayment      = "CREATE_PAYMENT"

	StatusExecute        = "STATUS_EXUCUTE"
	StatusSucess         = "STATUS_SUCCESS"
	StatusFailed         = "STATUS_FAILED"
	StatusRollbacked     = "STATUS_ROLLBACKED"
	StatusRollbackFailed = "STATUS_ROLLBACK_FAIL"
)

// PurchaseResult event
type BookingResult struct {
	CustomerID uint32
	BookingID  uint32
	Step       string
	Status     string
	Timestamp  time.Time
}
