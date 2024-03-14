package authentication

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/authentication"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

func validateAuthentication(e *pb.Authentication) error {
	if e.Username == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}

	if e.Password == "" {
		return status.Error(codes.InvalidArgument, "Password is required")
	}

	return nil
}

func validateUsername(username string) error {
	if username == "" {
		return status.Error(codes.InvalidArgument, "Username is required")
	}

	return nil
}

func validateCheckingOtpExpiration(e *pb.CheckingOtpExpirationRequest) error {
	if err := validateUsername(e.GetUsername()); err != nil {
		return err
	}
	if e.GetOtp() == "" {
		return status.Error(codes.InvalidArgument, "OTP is required")
	}

	return nil
}

func ValidEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
