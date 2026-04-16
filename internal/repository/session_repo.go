package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"id.diengs.backend/internal/entity"
)

type SessionRepo struct {
	Repository[entity.Session]
	Log *logrus.Logger
}

func NewSessionRepo(log *logrus.Logger) *SessionRepo {
	return &SessionRepo{
		Log: log,
	}
}

func (c *SessionRepo) FindByToken(db *gorm.DB, session *entity.Session, token string) error {
	return db.Where("token = ?", token).First(session).Error
}

func (r *SessionRepo) CountByUserId(db *gorm.DB, userId string) (int64, error) {
	var total int64
	err := db.Model(new(entity.Session)).Where("user_id = ?", userId).Count(&total).Error
	return total, err
}

func (r *SessionRepo) DeleteByUserId(db *gorm.DB, userId string) error {
	return db.Where("user_id = ?", userId).Delete(new(entity.Session)).Error
}

func (r *SessionRepo) DeleteByToken(db *gorm.DB, token string) error {
	return db.Where("token = ?", token).Delete(new(entity.Session)).Error
}
