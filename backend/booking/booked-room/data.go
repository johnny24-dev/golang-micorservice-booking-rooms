package booked_room

import (
	model "github.com/nekizz/final-project/backend/booking/model"
	pb "github.com/nekizz/final-project/backend/go-pbtype/bookedroom"
)

func prepareBookingToResponse(p *model.BookedRoom) *pb.BookedRoom {
	return &pb.BookedRoom{
		Id:       int32(p.ID),
		CheckIn:  p.CheckIn.String(),
		CheckOut: p.CheckOut.String(),
		Price:    p.Price,
		Quantity: uint32(p.Quantity),
	}
}
