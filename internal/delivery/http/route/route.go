package route

import (
	"id.diengs.backend/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App *fiber.App
	// Middleware
	AuthMiddleware       fiber.Handler
	AdminMiddleware      fiber.Handler
	HealthController     *http.HealthController
	AuthController       *http.AuthController
	ExperienceController *http.ExperienceController
}

func (c RouteConfig) Setup() {
	c.App.Get("/", c.HealthController.Check)
	c.App.Get("/api/health", c.HealthController.Check)
	c.SetupAuth()
	c.SetupExperience()
}

func (c RouteConfig) SetupAuth() {
	auth := c.App.Group("/api/auth")
	auth.Post("/send-otp", c.AuthController.SendOtp)
	auth.Post("/verify-otp", c.AuthController.VeriftOtp)
	auth.Post("/google", c.AuthController.AuthGoogle)
	auth.Delete("/_logout", c.AuthController.Logout)

	loggedRoute := auth.Group("/", c.AuthMiddleware)
	loggedRoute.Get("/_current", c.AuthController.Current)
}

func (c RouteConfig) SetupExperience() {
	experience := c.App.Group("/api/experiences")
	experience.Get("/", c.ExperienceController.Search)

	adminRoute := experience.Group("/", c.AuthMiddleware, c.AdminMiddleware)
	adminRoute.Post("/", c.ExperienceController.Create)
}
