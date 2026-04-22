package config

import (
	"gorm.io/gorm"
	"id.diengs.backend/internal/delivery/http"
	"id.diengs.backend/internal/delivery/http/route"
	"id.diengs.backend/internal/delivery/middleware"
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
	sessionRepo := repository.NewSessionRepo(cfg.Log)
	experienceRepo := repository.NewExperienceRepo(cfg.Log)
	experienceImageRepo := repository.NewExperienceImageRepo(cfg.Log)
	propertyRepo := repository.NewPropertyRepo(cfg.Log)
	hostProfileRepo := repository.NewHostProfileRepo(cfg.Log)

	// Use Case Config
	healthUseCase := usecase.NewHealthUseCase(cfg.Config)
	authUseCase := usecase.NewAuthUseCase(cfg.DB, cfg.Log, cfg.Validate, cfg.Mail, userRepo, emailOtpRepo, sessionRepo, cfg.Config)
	experienceUseCase := usecase.NewExperienceUseCase(cfg.DB, cfg.Log, cfg.Validate, experienceRepo, experienceImageRepo)
	propertyUseCase := usecase.NewPropertyUseCase(cfg.DB, cfg.Log, cfg.Validate, propertyRepo, hostProfileRepo)

	// Controller Config
	healthController := http.NewHealthController(healthUseCase, cfg.Log)
	authController := http.NewAuthController(authUseCase, cfg.Log)
	experienceController := http.NewExperienceController(experienceUseCase, cfg.Log)
	propertyCotroller := http.NewPropertyController(cfg.Log, propertyUseCase)

	// setup middleware
	authMiddleware := middleware.NewAuth(authUseCase)
	adminMiddleware := middleware.NewAdmin()

	route.RouteConfig{
		App:                  cfg.App,
		AuthMiddleware:       authMiddleware,
		AdminMiddleware:      adminMiddleware,
		HealthController:     healthController,
		AuthController:       authController,
		ExperienceController: experienceController,
		PropertyController:   propertyCotroller,
	}.Setup()
}
