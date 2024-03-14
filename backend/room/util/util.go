package util

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

//func DecodeCreateBookingCmd(payload message.Payload) (*modelBK.Booking, *pbB.Booking, error) {
//	var cmd pbB.CreateBookingCmd
//
//	if err := json.Unmarshal(payload, &cmd); err != nil {
//		return nil, nil, err
//	}
//
//	bookingID := cmd.BookingId
//	bookDate, _ := time.Parse(time.RFC3339, cmd.Booking.GetBookDate())
//
//	return &modelBK.Booking{
//		Model: gorm.Model{
//			ID: uint(bookingID),
//		},
//		BookDate:   &bookDate,
//		Note:       cmd.Booking.Note,
//		CustomerID: uint(cmd.Booking.CustomerId),
//	}, cmd.Booking, nil
//}

func Time2pbTimestamp(now time.Time) *timestamp.Timestamp {
	s := int64(now.Unix())
	n := int32(now.Nanosecond())
	return &timestamp.Timestamp{Seconds: s, Nanos: n}
}
