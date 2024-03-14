package hotel

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateOne(e *pb.OneHotelRequest) error {
	if e.GetId() == "" {
		return status.Error(codes.InvalidArgument, "HotelID is required")
	}

	return nil
}

func validateUpdate(e *pb.Hotel) error {
	return validateCreate(e)
}

func validateCreate(e *pb.Hotel) error {
	if e.GetName() == "" {
		return status.Error(codes.InvalidArgument, "Name is required")
	}

	if e.GetStarLevel() > 5 || e.GetStarLevel() < 0 {
		return status.Error(codes.InvalidArgument, "Star level must be from 0 to 5")
	}

	if e.GetAddress() == nil {
		return status.Error(codes.InvalidArgument, "Address is required")
	}

	if e.GetListImage() == nil {
		return status.Error(codes.InvalidArgument, "Must be upload at least 1 image")
	}

	return nil
}

func validateList(page, limit uint32) error {
	if page <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid Page")
	}

	if limit <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid Limit")
	}

	return nil
}
