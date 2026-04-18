package repository

import (
	"github.com/sirupsen/logrus"
	"id.diengs.backend/internal/entity"
)

type HostProfileRepo struct {
	Repository[entity.HostProfile]
	Log *logrus.Logger
}

func NewHostProfileRepo(log *logrus.Logger) *HostProfileRepo {
	return &HostProfileRepo{
		Log: log,
	}
}
