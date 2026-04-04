package usecase

import (
	"id.diengs.backend/internal/entity"
	"id.diengs.backend/internal/model"

	"github.com/spf13/viper"
)

type HealthUseCase struct {
	Config *viper.Viper
}

func NewHealthUseCase(config *viper.Viper) *HealthUseCase {
	return &HealthUseCase{Config: config}
}

func (u *HealthUseCase) Check() model.HealthResponse {
	health := entity.Health{
		Name:    u.Config.GetString("app.name"),
		Version: "starter",
		Status:  "UP",
	}

	return model.HealthResponse{
		Name:    health.Name,
		Version: health.Version,
		Status:  health.Status,
	}
}
