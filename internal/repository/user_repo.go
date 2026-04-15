package repository

import (
	"github.com/sirupsen/logrus"
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
