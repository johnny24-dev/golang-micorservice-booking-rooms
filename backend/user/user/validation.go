package user

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

func validateOne(e *pb.OneUserRequest) error {
	if e.GetId() == 0 {
		return status.Error(codes.InvalidArgument, "UserID is required")
	}

	return nil
}

func ValidateCreateUser(e *pb.User) error {
	if e.Name == "" {
		return status.Error(codes.InvalidArgument, "Name is required")
	}
	if e.Gender == "" {
		return status.Error(codes.InvalidArgument, "Gender is required")
	}
	if e.Email == "" {
		return status.Error(codes.InvalidArgument, "Email is required")
	}
	if e.GetAccountId() == 0 {
		return status.Error(codes.InvalidArgument, "AccountID is required")
	}

	return nil
}

func validateUpdate(e *pb.User) error {
	//if e.Email == "" && e.Phone == "" && e.Address == "" && e.Avatar == "" && e.Token == "" {
	//	return status.Error(codes.InvalidArgument, "Change field can't be empty")
	//}
	//if e.Phone != "" && !ValidPhone(e.Phone) {
	//	return status.Error(codes.InvalidArgument, "Wrong type of phone")
	//}
	//if e.Email != "" && !ValidEmail(e.Email) {
	//	return status.Error(codes.InvalidArgument, "Wrong type of email")
	//}

	return nil
}

func ValidPhone(textCheck string) bool {
	re := regexp.MustCompile(`^((03|07|09|08|01[2|6|8|9])+([0-9]{8})|(843|847|849|841[2|6|8|9])+([0-9]{8}))$`)
	return re.MatchString(textCheck)
}

func ValidEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
