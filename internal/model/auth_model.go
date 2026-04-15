package model

type AuthSendOtpReq struct {
	Email string `json:"email" validate:"required"`
}

type AuthVerifyOtpRequest struct {
	Email string `json:"email" validate:"required"`
	Otp   string `json:"otp" validate:"required"`
}
