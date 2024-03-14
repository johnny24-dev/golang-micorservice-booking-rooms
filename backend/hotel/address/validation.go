package address

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/address"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateUpdate(e *pb.Address) error {
	if e == nil {
		return status.Error(codes.InvalidArgument, "Address is required")
	}
	return nil
}

func validateOne(e *pb.OneAddressRequest) error {
	if e.GetId() == 0 {
		return status.Error(codes.InvalidArgument, "AddressID is required")
	}
	return nil
}
