package booking

import (
	"context"
	"errors"
	"fmt"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	booked_room "github.com/nekizz/final-project/backend/booking/booked-room"
	"github.com/nekizz/final-project/backend/booking/logger"
	model "github.com/nekizz/final-project/backend/booking/model"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	pbn "github.com/nekizz/final-project/backend/go-pbtype/notification"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
	"time"
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

func (s *Service) ConfirmBooking(ctx context.Context, r *pb.Booking) (*pb.Booking, error) {
	if err := validateCreate(r); err != nil {
		return nil, err
	}

	tx := s.db.Begin()

	payment := &model.Payment{
		Model:           gorm.Model{},
		PaypalPaymentID: r.Payment.PaypalPaymentId,
		Status:          r.Payment.Status,
		BookingID:       uint(r.GetId()),
	}
	queryPayment := tx.Model(model.Payment{}).Create(payment)
	if err := queryPayment.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	//update booking
	queryBooking := tx.Model(model.Booking{}).Where("id = ?", r.GetId()).Updates(map[string]interface{}{"status": "approved"})
	if err := queryBooking.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return r, nil
}

func (s *Service) CreateBooking(ctx context.Context, r *pb.BookingV2) (*pb.BookingV2, error) {
	if err := validateCreateV2(r); err != nil {
		return nil, err
	}

	checkInPb := r.GetBookedroom()[0].GetCheckIn()
	checkOutPb := r.GetBookedroom()[0].GetCheckOut()

	checkIn, _ := time.Parse("2006-01-02 15:04:05.000 -07", checkInPb)
	checkOut, _ := time.Parse("2006-01-02 15:04:05.000 -07", checkOutPb)

	for _, v := range r.GetBookedroom() {
		var quantityRoom int64
		var mapping map[string]interface{}

		queryR := s.db.Table("rooms").Select("quantity").Where("id = ?", v.GetRoomId()).Scan(&quantityRoom)
		if err := queryR.Error; err != nil {
			return nil, err
		}

		queryBR := s.db.Table("booked_rooms").Select("SUM(quantity) as countBookedRoom").Where("room_id = ? AND check_in >= ? AND check_out <= ?", v.GetRoomId(), checkIn, checkOut).Find(&mapping)
		if err := queryBR.Error; err != nil {
			return nil, err
		}

		count, _ := mapping["countBookedRoom"].(string)
		countBookedRoom, _ := strconv.Atoi(count)

		if int(v.GetQuantity())+countBookedRoom > int(quantityRoom) {
			return nil, errors.New("Not enough quantity of room_id: " + strconv.Itoa(int(v.GetRoomId())))
		}
	}

	booking, err := NewRepository(s.db).CreatOne(prepareBookingV2ToRequest(r))
	if nil != err {
		return nil, err
	}

	var listBookedRoom []*model.BookedRoom

	for _, val := range r.GetBookedroom() {
		checkIn, err := time.Parse("2006-01-02 15:04:05.000 -07", val.GetCheckIn())
		if err != nil {
			fmt.Println(err)
		}
		checkOut, err := time.Parse("2006-01-02 15:04:05.000 -07", val.GetCheckOut())
		if err != nil {
			fmt.Println(err)
		}
		bookedRoom := &model.BookedRoom{
			Model:     gorm.Model{},
			CheckIn:   &checkIn,
			CheckOut:  &checkOut,
			Price:     val.GetPrice(),
			IsCheckIn: "false",
			BookingID: booking.ID,
			Quantity:  uint(val.Quantity),
			RoomID:    uint(val.GetRoomId()),
		}
		listBookedRoom = append(listBookedRoom, bookedRoom)
	}
	queryBookedRoom := s.db.Model(model.BookedRoom{}).Create(listBookedRoom)
	if err := queryBookedRoom.Error; err != nil {
		return nil, err
	}

	return prepareBookingV2ToResponse(booking), nil
}

func (s *Service) ListBookingOfCustomer(ctx context.Context, r *pb.ListCustomerBookingRequest) (*pb.ListCustomerBookingResponse, error) {
	if err := validateListCustomerBooking(r); nil != err {
		return nil, err
	}

	var list []*pb.Booking

	bookings, count, err := NewRepository(s.db).ListCustomerBooking(r)
	if nil != err {
		return nil, err
	}

	for _, booking := range bookings {
		listBookedRoom, total, hotel, err := booked_room.GetIDRoomByBookingID(int(booking.ID), s.db)
		if err != nil {
			fmt.Println(err)
		}
		bookingPb := prepareBookingToResponse(booking)
		bookingPb.Bookedroom = listBookedRoom
		bookingPb.TotalPrice = total
		bookingPb.Hotel = hotel
		list = append(list, bookingPb)
	}

	return &pb.ListCustomerBookingResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s *Service) ListBookingOfHost(ctx context.Context, r *pb.ListHostBookingRequest) (*pb.ListHostBookingResponse, error) {
	if err := validateListHostBooking(r); nil != err {
		return nil, err
	}

	var list []*pb.Booking

	bookings, count, err := NewRepository(s.db).ListHostBooking(r)
	if nil != err {
		return nil, err
	}

	for _, booking := range bookings {
		listBookedRoom, total, hotel, err := booked_room.GetIDRoomByBookingID(int(booking.ID), s.db)
		if err != nil {
			return nil, err
		}
		bookingPb := prepareBookingToResponse(booking)
		bookingPb.Bookedroom = listBookedRoom
		bookingPb.TotalPrice = total
		bookingPb.Hotel = hotel
		list = append(list, bookingPb)
	}

	return &pb.ListHostBookingResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s *Service) Cancel(ctx context.Context, r *pb.OneBookingRequest) (*pbn.Notification, error) {
	if err := validateOne(r); nil != err {
		return &pbn.Notification{
			Status:  "Bad request",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}, nil
	}

	err := NewRepository(s.db).UpdateStatusHotel(int(r.GetId()), "cancelled")
	if nil != err {
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}, nil
	}

	return &pbn.Notification{
		Status:  "OK",
		Message: "Cancel boooking successfully!",
		Code:    http.StatusOK,
	}, nil
}

func (s *Service) ChangeCompletedBooking(ctx context.Context, r *pb.OneBookingRequest) (*pbn.Notification, error) {
	if err := validateOne(r); nil != err {
		return &pbn.Notification{
			Status:  "Bad request",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}, nil
	}

	err := NewRepository(s.db).UpdateStatusHotel(int(r.GetId()), "completed")
	if nil != err {
		return &pbn.Notification{
			Status:  "Internal Error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}, nil
	}

	return &pbn.Notification{
		Status:  "OK",
		Message: "Complete boooking successfully!",
		Code:    http.StatusOK,
	}, nil
}

func AutoCheckCompletedBooking(ctx context.Context, db *gorm.DB) error {
	logger.Log.WithFields(logrus.Fields{
		"service":  "Booking Service",
		"function": "AutoCheckApprovedBooking",
	}).Println("Start running cronjob!")
	hourRun := 0
	tick := time.NewTicker(time.Hour)
	defer tick.Stop()
	for {
		logger.Log.WithFields(logrus.Fields{
			"service":  "Booking Service",
			"function": "AutoCheckApprovedBooking",
		}).Println("Running hour: ", hourRun) // 1h lan quet lai db
		select {
		case <-tick.C:
			timeNow := time.Now().Add(7 * time.Hour)
			query := db.Model(model.Booking{}).Where("status = ? AND end_date < ?", "completed", timeNow).Updates(map[string]interface{}{"status": "completed"})
			if err := query.Error; nil != err {
				logger.Log.WithFields(logrus.Fields{
					"service":  "Booking Service",
					"function": "AutoCheckApprovedBooking",
				}).WithError(err).Error("update status approved fail")
			}
			hourRun++
		case <-ctx.Done():
			return nil
		}
	}

	logger.Log.WithFields(logrus.Fields{
		"service":  "Booking Service",
		"function": "AutoCheckApprovedBooking",
	}).Println("Stop running cronjob!")

	return nil
}
