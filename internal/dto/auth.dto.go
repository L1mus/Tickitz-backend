package dto

type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	AgreeTerms bool   `json:"agree_terms" binding:"required"`
}

type RegisterResponse struct {
	Message   string `json:"message"`
	Email     string `json:"email"`
	Is_Active bool   `json:"is_active"`
}

type ActivationRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	Message string      `json:"message"`
	Token   string      `json:"token"`
	User    UserDetails `json:"user"`
}

type UserDetails struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
type VerifyResetOTPReq struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,len=6"`
}
type ResetPasswordReq struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

type ResendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}
