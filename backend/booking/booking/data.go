package booking

import (
	model "github.com/nekizz/final-project/backend/booking/model"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"gorm.io/gorm"
	"time"
)

func prepareBookingV2ToRequest(p *pb.BookingV2) *model.BookingV2 {
	startDate, _ := time.Parse("2006-01-02 15:04:05.000 -07", p.GetStartDate())
	endDate, _ := time.Parse("2006-01-02 15:04:05.000 -07", p.GetEndDate())

	return &model.BookingV2{
		Model:      gorm.Model{},
		Status:     "pending",
		Note:       p.GetNote(),
		StartDate:  &startDate,
		EndDate:    &endDate,
		CustomerID: uint(p.GetCustomerId()),
	}
}

func prepareBookingV2ToResponse(p *model.BookingV2) *pb.BookingV2 {
	return &pb.BookingV2{
		Id:         int32(p.ID),
		Note:       p.Note,
		StartDate:  p.StartDate.String(),
		EndDate:    p.EndDate.String(),
		CustomerId: int32(p.CustomerID),
	}
}

func prepareBookingToResponse(p *model.Booking) *pb.Booking {
	return &pb.Booking{
		Id:         int32(p.ID),
		Note:       p.Note,
		Status:     p.Status,
		StartDate:  p.StartDate.String(),
		EndDate:    p.EndDate.String(),
		CustomerId: int32(p.CustomerID),
	}
}
