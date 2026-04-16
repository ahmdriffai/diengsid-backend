package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"id.diengs.backend/internal/entity"
)

type UserRepo struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepo(log *logrus.Logger) *UserRepo {
	return &UserRepo{
		Log: log,
	}
}

func (r *UserRepo) FindByEmail(db *gorm.DB, entity *entity.User, email string) error {
	return db.Where("email = ?", email).Take(entity).Error
}
