package service

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateOne(e *pb.OneServiceRequest) error {
	if e.GetId() == "" {
		return status.Error(codes.InvalidArgument, "UserID is required")
	}

	return nil
}

func validateUpdate(e *pb.Service) error {
	if e.GetId() == 0 {
		return status.Error(codes.InvalidArgument, "Service ID is required")
	}

	return validateCreate(e)
}

func validateCreate(e *pb.Service) error {
	if e.GetName() == "" {
		return status.Error(codes.InvalidArgument, "Name is required")
	}

	return nil
}

func validateList(e *pb.ListServiceRequest) error {
	if e.GetPage() <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid Page")
	}

	if e.GetLimit() <= 0 {
		return status.Error(codes.InvalidArgument, "Invalid Limit")
	}

	return nil
}
