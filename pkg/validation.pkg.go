package pkg

import (
	"regexp"
	"strings"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func parseValidationError(err error) string {
	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "Email") && strings.Contains(errStr, "email"):
		return apperror.ErrInvalidEmail.Error()
	case strings.Contains(errStr, "Password") && strings.Contains(errStr, "min=8"):
		return apperror.ErrInvalidPassword.Error()
	case strings.Contains(errStr, "required"):
		return "required field is missing"
	default:
		return "invalid input data"
	}
}
