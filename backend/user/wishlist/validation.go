package wishlist

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/wishlist"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateCreate(r *pb.Wishlist) error {
	if r.GetHotelId() == 0 {
		return status.Error(codes.InvalidArgument, "HotelId is required")
	}
	if r.GetCustomerId() == 0 {
		return status.Error(codes.InvalidArgument, "CustomerID is required")
	}

	return nil
}

func validateDelete(r *pb.Wishlist) error {
	if r.GetHotelId() == 0 {
		return status.Error(codes.InvalidArgument, "HotelId is required")
	}
	if r.GetCustomerId() == 0 {
		return status.Error(codes.InvalidArgument, "CustomerID is required")
	}

	return nil
}

func validateList(r *pb.ListWishlistRequest) error {
	if r.GetCustomerId() == 0 {
		return status.Error(codes.InvalidArgument, "CustomerID is required")
	}

	return nil
}
