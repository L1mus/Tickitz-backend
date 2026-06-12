package apperror

import "errors"

var (
	ErrInvalidEmail         = errors.New("invalid email format")
	ErrInvalidPassword      = errors.New("invalid password format min 8 characters")
	ErrUserNotFound         = errors.New("user not found")
	ErrInternalServer       = errors.New("internal server error")
	ErrEmailRegistered      = errors.New("email already registered")
	ErrOTPInvalid           = errors.New("invalid otp code")
	ErrOTPExpired           = errors.New("otp expired or not found")
	ErrInvalidCredentials   = errors.New("invalid email or password")
	ErrCredentialsRequired  = errors.New("email and password are required")
	MovieNotFound           = errors.New("movie not found")
	MovieGenresNotFound     = errors.New("genre not found")
	MovieCastsNotFound      = errors.New("cast not found")
	InvalidSeatsInput       = errors.New("seats is required")
	InvalidQuantity         = errors.New("quantity is not the same as the number of seats ordered")
	SeatsUnavailable        = errors.New("seat already taken")
	ErrAccountNotActivated  = errors.New("ACCOUNT_NOT_ACTIVATED")
	ErrAccountAlreadyActive = errors.New("account is already activated")
	ErrResetTokenExpired    = errors.New("unauthorized or session expired, please verify OTP again")
	ErrMissingAuthHeader    = errors.New("authorization header is required")
	ErrTokenRevoked         = errors.New("token has been revoked or logged out")
)
