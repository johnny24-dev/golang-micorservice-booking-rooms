package customer

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/customer"
	"github.com/nekizz/final-project/backend/user/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateOne(e *pb.OneCustomerRequest) error {
	if e.GetId() == "" {
		return status.Error(codes.InvalidArgument, "UserID is required")
	}

	return nil
}

func validateCreateCustomer(e *pb.Customer) error {
	if e.User == nil {
		return status.Error(codes.InvalidArgument, "User is required")
	}

	if err := user.ValidateCreateUser(e.User); err != nil {
		return err
	}
	return nil
}

func validateUpdate(e *pb.Customer) error {
	if e == nil {
		return status.Error(codes.InvalidArgument, "Customer is required")
	}

	if e.User == nil {
		return status.Error(codes.InvalidArgument, "User is required")
	}
	return nil
}
