package booking

import (
	"errors"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateCreate(e *pb.Booking) error {
	if e.GetPayment() == nil {
		return status.Error(codes.InvalidArgument, "Payment is required")
	}

	if e.GetId() == 0 {
		return status.Error(codes.InvalidArgument, "BookingID is required")
	}

	return nil
}

func validateCreateV2(e *pb.BookingV2) error {
	if len(e.GetBookedroom()) == 0 {
		return status.Error(codes.InvalidArgument, "Booked Room is required")
	}

	if len(e.GetBookedroom()) > 10 {
		return status.Error(codes.InvalidArgument, "Maximum Booking 10 hotels")
	}

	if e.GetStartDate() == "" {
		return status.Error(codes.InvalidArgument, "Start date is required")
	}

	if e.GetEndDate() == "" {
		return status.Error(codes.InvalidArgument, "End date is required")
	}

	if e.GetCustomerId() <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid customer_id")
	}

	return nil
}

func validateListCustomerBooking(e *pb.ListCustomerBookingRequest) error {
	if e.GetCustomerId() == 0 {
		return status.Error(codes.InvalidArgument, "CustomerID is required")
	}

	return nil
}

func validateListHostBooking(e *pb.ListHostBookingRequest) error {
	if e.GetHostId() == 0 {
		return status.Error(codes.InvalidArgument, "HostID is required")
	}

	return nil
}

func validateOne(r *pb.OneBookingRequest) error {
	if r.GetId() == 0 {
		return errors.New("Booking Id is required")
	}

	return nil
}
