package config

import (
	"gorm.io/gorm"
	"id.diengs.backend/internal/delivery/http"
	"id.diengs.backend/internal/delivery/http/route"
	"id.diengs.backend/internal/pkg"
	"id.diengs.backend/internal/repository"
	"id.diengs.backend/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Mail     *pkg.Mail
}

func Bootstrap(cfg *BootstrapConfig) {
	// Repository Config
	userRepo := repository.NewUserRepo(cfg.Log)
	emailOtpRepo := repository.NewEmailOtpRepo(cfg.Log)

	// Use Case Config
	healthUseCase := usecase.NewHealthUseCase(cfg.Config)
	authUseCase := usecase.NewAuthUseCase(cfg.DB, cfg.Log, cfg.Validate, cfg.Mail, userRepo, emailOtpRepo)

	// Controller Config
	healthController := http.NewHealthController(healthUseCase, cfg.Log)
	authController := http.NewAuthController(authUseCase, cfg.Log)

	route.RouteConfig{
		App:              cfg.App,
		HealthController: healthController,
		AuthController:   authController,
	}.Setup()
}
