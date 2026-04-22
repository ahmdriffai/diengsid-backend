package model

import "id.diengs.backend/internal/entity"

type HostProfileResponse struct {
	ID                string `json:"id"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Address           string `json:"address"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number"`
	KTPNumber         string `json:"ktp_number"`
	Bio               string `json:"bio"`
	CreatedAt         int64  `json:"created_at"`
	UpdatedAt         int64  `json:"updated_at"`
}

type HostCreateRequest struct {
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Address           string `json:"address"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number"`
	KTPNumber         string `json:"ktp_number"`
	Bio               string `json:"bio"`
}

func HostToResponse(host *entity.HostProfile) *HostProfileResponse {
	return &HostProfileResponse{
		ID:                host.ID,
		PhoneNumber:       host.PhoneNumber,
		ProfilePictureURL: host.ProfilePictureURL,
		Address:           host.Address,
		BankAccountName:   host.BankAccountName,
		BankAccountNumber: host.BankAccountNumber,
		KTPNumber:         host.KTPNumber,
		Bio:               host.Bio,
		CreatedAt:         host.CreatedAt,
		UpdatedAt:         host.UpdatedAt,
	}
}
