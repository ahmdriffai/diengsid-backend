package route

import (
	"id.diengs.backend/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App              *fiber.App
	HealthController *http.HealthController
	AuthController   *http.AuthController
	AuthMiddleware   fiber.Handler
}

func (c RouteConfig) Setup() {
	c.App.Get("/", c.HealthController.Check)
	c.App.Get("/api/health", c.HealthController.Check)
	c.SetupAuth()
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
