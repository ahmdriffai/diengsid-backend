package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Experience struct {
	ID             string   `gorm:"column:id;primaryKey"`
	ExperienceType string   `gorm:"column:experience_type;not null"`
	Title          string   `gorm:"column:title;not null"`
	Address        string   `gorm:"column:address;not null"`
	Description    string   `gorm:"column:description;type:text;not null"`
	ThumbnailURL   *string  `gorm:"column:thumbnail_url"`
	Lat            *float64 `gorm:"column:lat"`
	Lng            *float64 `gorm:"column:lng"`
	BasePrice      float64  `gorm:"column:base_price;not null"`

	Images []ExperienceImage `gorm:"foreignKey:ExperienceID;constraint:OnDelete:CASCADE"`

	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

func (Experience) TableName() string {
	return "experiences"
}

func (e *Experience) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	e.CreatedAt = time.Now().UnixMilli()
	e.UpdatedAt = time.Now().UnixMilli()
	return nil
}

type ExperienceImage struct {
	ID           string `gorm:"column:id;primaryKey"`
	ExperienceID string `gorm:"column:experience_id;not null"`
	ImageURL     string `gorm:"column:image_url;not null"`
	IsPrimary    bool   `gorm:"column:is_primary;default:false"`

	Experience Experience `gorm:"foreignKey:ExperienceID;references:ID;constraint:OnDelete:CASCADE"`

	CreatedAt int64 `gorm:"column:created_at"`
	UpdatedAt int64 `gorm:"column:updated_at"`
}

func (ExperienceImage) TableName() string {
	return "experience_images"
}

func (e *ExperienceImage) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	e.CreatedAt = time.Now().UnixMilli()
	e.UpdatedAt = time.Now().UnixMilli()
	return nil
}
