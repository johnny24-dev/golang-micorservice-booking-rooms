package account

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

func validateOne(e *pb.OneAccountRequest) error {
	if e.GetId() == "" {
		return status.Error(codes.InvalidArgument, "AccountID is required")
	}

	return nil
}

func ValidateCreateAccount(e *pb.Account) error {
	if e.Username == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}
	if e.Password == "" {
		return status.Error(codes.InvalidArgument, "Password is required")
	}

	return nil
}

func validateVerificationEmail(e *pb.OneAccountRequest) error {
	if e.Username == "" {
		return status.Error(codes.InvalidArgument, "Email is required")
	}

	if !ValidEmail(e.GetUsername()) {
		return status.Error(codes.InvalidArgument, "Wrong type of email")
	}

	return nil
}

func validateCheckingOtpExpiration(e *pb.CheckingOtpExpirationRequest) error {
	if e.GetOtp() == "" {
		return status.Error(codes.InvalidArgument, "OTP is required")
	}
	if e.GetUsername() == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}

	return nil
}

func validateLogin(e *pb.Account) error {
	if e.Username == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}
	if e.Password == "" {
		return status.Error(codes.InvalidArgument, "Password is required")
	}

	return nil
}

func validateCheckExsit(username string) error {
	if username == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}

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
