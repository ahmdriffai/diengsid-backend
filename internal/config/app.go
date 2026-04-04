package config

import (
	"id.diengs.backend/internal/delivery/http"
	"id.diengs.backend/internal/delivery/http/route"
	"id.diengs.backend/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(cfg *BootstrapConfig) {
	healthUseCase := usecase.NewHealthUseCase(cfg.Config)
	healthController := http.NewHealthController(healthUseCase, cfg.Log)

	route.RouteConfig{
		App:              cfg.App,
		HealthController: healthController,
	}.Setup()
}
