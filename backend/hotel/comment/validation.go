package comment

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/comment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateCreate(e *pb.Comment) error {
	if e.GetText() == "" {
		return status.Error(codes.InvalidArgument, "Text is required")
	}

	if e.GetCustomerId() == 0 {
		return status.Error(codes.InvalidArgument, "CustomerID is required")
	}

	if e.GetHotelId() == 0 {
		return status.Error(codes.InvalidArgument, "HotelID is required")
	}

	if e.GetRate() < 0 || e.GetRate() > 10 {
		return status.Error(codes.InvalidArgument, "Rate is invalid type")
	}

	return nil
}

func validateDeleteC(e *pb.OneCommentRequest) error {
	if e.GetId() == 0 {
		return status.Error(codes.InvalidArgument, "CommentID is required")
	}

	return nil
}

func validateList(r *pb.ListCommentRequest) error {
	if r.GetHotelId() == 0 {
		return status.Error(codes.InvalidArgument, "HotelID is required")
	}

	if r.GetPage() <= 0 {
		return status.Error(codes.InvalidArgument, "HotelID is required")
	}

	if r.GetLimit() <= 0 {
		return status.Error(codes.InvalidArgument, "HotelID is required")
	}

	if r.GetSortAsc() != 0 && r.GetSortAsc() != 1 {
		return status.Error(codes.InvalidArgument, "Sort_number is invalid")
	}
	return nil
}
