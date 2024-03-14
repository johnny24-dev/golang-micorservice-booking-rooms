package room

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/room"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func validateOne(e *pb.OneRoomRequest) error {
	if e.GetId() == "" {
		return status.Error(codes.InvalidArgument, "RoomID is required")
	}

	return nil
}

func validateUpdate(e *pb.Room) error {
	if e.GetName() == "" {
		return status.Error(codes.InvalidArgument, "Name is required")
	}
	if e.GetQuantity() == 0 {
		return status.Error(codes.InvalidArgument, "Quantity is required")
	}

	return nil
}

func validateCreate(e *pb.Room) error {
	if e.GetName() == "" {
		return status.Error(codes.InvalidArgument, "Name is required")
	}
	if e.GetHotelId() == 0 {
		return status.Error(codes.InvalidArgument, "HotelID is required")
	}
	if e.GetQuantity() == 0 {
		return status.Error(codes.InvalidArgument, "Quantity is required")
	}

	return nil
}

func validateList(e *pb.ListRoomRequest) error {
	var field = strings.Split(e.GetSearchField(), ",")

	if e.GetPage() <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid Page")
	}

	if e.GetLimit() <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid Limit")
	}

	if e.GetSearchField() == "" && e.GetSearchValue() == "" {
		return nil
	}

	for i, _ := range field {
		var newField = strings.ToLower(strings.TrimSpace(field[i]))
		if newField != "name" && newField != "dob" && newField != "email" && newField != "gender" && newField != "role" && newField != "" {
			return status.Error(codes.InvalidArgument, "Invalid SearchFields")
		}
	}

	return nil
}
