package repository

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"id.diengs.backend/internal/entity"
)

type EmailOtpRepo struct {
	Repository[entity.EmailOtp]
	Log *logrus.Logger
}

func NewEmailOtpRepo(log *logrus.Logger) *EmailOtpRepo {
	return &EmailOtpRepo{
		Log: log,
	}
}

func (r *EmailOtpRepo) FindActiveAndEmail(db *gorm.DB, entity *entity.EmailOtp, email string) error {
	return db.Where("email = ? AND is_used = ? AND expired_at > ?", email, false, time.Now().UnixMilli()).
		Order("created_at DESC").
		Limit(1).
		Take(entity).Error
}
