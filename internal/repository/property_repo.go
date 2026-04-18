package repository

import (
	"github.com/sirupsen/logrus"
	"id.diengs.backend/internal/entity"
)

type PropertyRepo struct {
	Repository[entity.Property]
	Log *logrus.Logger
}

func NewPropertyRepo(log *logrus.Logger) *PropertyRepo {
	return &PropertyRepo{
		Log: log,
	}
}
