package config

import (
	"id.diengs.backend/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func NewFiber(cfg *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      cfg.GetString("app.name"),
		Prefork:      cfg.GetBool("app.prefork"),
		ErrorHandler: NewErrorHandler(),
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
	}))

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if fiberErr, ok := err.(*fiber.Error); ok {
			code = fiberErr.Code
		}

		return ctx.Status(code).JSON(model.WebResponse[any]{
			Success: false,
			Message: err.Error(),
		})
	}
}
