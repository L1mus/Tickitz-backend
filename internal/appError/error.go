package apperror

import "errors"

var (
	ErrInvalidEmail        = errors.New("invalid email format")
	ErrInvalidPassword     = errors.New("invalid password format min 8 characters")
	ErrUserNotFound        = errors.New("user not found")
	ErrInternalServer      = errors.New("internal server error")
	ErrEmailRegistered     = errors.New("email already registered")
	ErrOTPInvalid          = errors.New("invalid otp code")
	ErrOTPExpired          = errors.New("otp expired or not found")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrCredentialsRequired = errors.New("email and password are required")
)
