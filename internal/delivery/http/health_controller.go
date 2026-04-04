package http

import (
	"id.diengs.backend/internal/model"
	"id.diengs.backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type HealthController struct {
	UseCase *usecase.HealthUseCase
	Log     *logrus.Logger
}

func NewHealthController(useCase *usecase.HealthUseCase, log *logrus.Logger) *HealthController {
	return &HealthController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *HealthController) Check(ctx *fiber.Ctx) error {
	result := c.UseCase.Check()

	return ctx.JSON(model.WebResponse[model.HealthResponse]{
		Success: true,
		Message: "ok",
		Data:    result,
	})
}
