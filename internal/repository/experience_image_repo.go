package repository

import (
	"github.com/sirupsen/logrus"
	"id.diengs.backend/internal/entity"
)

type ExperienceImageRepo struct {
	Repository[entity.ExperienceImage]
	Log *logrus.Logger
}

func NewExperienceImageRepo(log *logrus.Logger) *ExperienceImageRepo {
	return &ExperienceImageRepo{
		Log: log,
	}
}
