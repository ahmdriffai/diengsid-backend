package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Property struct {
	ID           string `gorm:"column:id;primaryKey"`
	HostID       string `gorm:"column:host_id;not null"`
	ExperienceID string `gorm:"column:experience_id;not null"`
	PropertyType string `gorm:"column:property_type;default:homestay"`
	BookingType  string `gorm:"column:booking_type"`

	Host       HostProfile `gorm:"foreignKey:HostID;references:ID;constraint:OnDelete:CASCADE"`
	Experience Experience  `gorm:"foreignKey:ExperienceID;references:ID;constraint:OnDelete:CASCADE"`

	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

func (Property) TableName() string {
	return "properties"
}

func (p *Property) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	p.CreatedAt = time.Now().UnixMilli()
	p.UpdatedAt = time.Now().UnixMilli()
	return nil
}

type HostProfile struct {
	ID                string `gorm:"column:id;primaryKey"`
	PhoneNumber       string `gorm:"column:phone_number"`
	ProfilePictureURL string `gorm:"column:profile_picture_url"`
	Address           string `gorm:"column:address"`
	BankAccountName   string `gorm:"column:bank_account_name"`
	BankAccountNumber string `gorm:"column:bank_account_number"`
	KTPNumber         string `gorm:"column:ktp_number"`
	Bio               string `gorm:"column:bio"`

	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

func (HostProfile) TableName() string {
	return "host_profiles"
}

func (h *HostProfile) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.NewString()
	h.CreatedAt = time.Now().UnixMilli()
	h.UpdatedAt = time.Now().UnixMilli()
	return nil
}

type Rentable struct {
	ID string `gorm:"column:id;primaryKey"`
}
