package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/sony/sonyflake"
	"time"
)

type IDGenerator interface {
	NextID() (uint64, error)
}

func NewSonyFlake() (IDGenerator, error) {
	var st sonyflake.Settings
	sf := sonyflake.NewSonyflake(st)
	if sf == nil {
		return nil, errors.New("sonyflake not created")
	}
	return sf, nil
}

func PublishMessage(publisher *kafka.Publisher, task interface{}, topic string) error {
	payload, err := json.Marshal(task)
	if err != nil {
		return err
	}

	err = publisher.Publish(topic, &message.Message{UUID: watermill.NewUUID(), Payload: payload})
	if err != nil {
		return err
	}

	return nil
}

func Time2pbTimestamp(now time.Time) *timestamp.Timestamp {
	s := int64(now.Unix())
	n := int32(now.Nanosecond())
	return &timestamp.Timestamp{Seconds: s, Nanos: n}
}

func BuildQueryHostBooking(hostId int) string {
	sqlFilter := fmt.Sprintf("SELECT DISTINCT bookings.id FROM doandb.booked_rooms inner join bookings on booked_rooms.booking_id = bookings.id inner join rooms on booked_rooms.room_id = rooms.id inner join hotels on hotels.id = rooms.hotel_id and hotels.customer_id = %v", hostId)

	sqlQuery := fmt.Sprintf("id IN (%s)", sqlFilter)

	return sqlQuery
}
