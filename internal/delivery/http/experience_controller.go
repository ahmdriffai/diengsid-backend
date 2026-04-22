package http

import (
	"math"

	"id.diengs.backend/internal/model"
	"id.diengs.backend/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ExperienceController struct {
	ExperienceUseCase *usecase.ExperienceUseCase
	Log               *logrus.Logger
}

func NewExperienceController(experienceUseCase *usecase.ExperienceUseCase, log *logrus.Logger) *ExperienceController {
	return &ExperienceController{
		ExperienceUseCase: experienceUseCase,
		Log:               log,
	}
}

// create experience
func (c *ExperienceController) Create(ctx *fiber.Ctx) error {
	request := new(model.ExperienceCreateRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("failed to parse request body")
		return fiber.ErrBadRequest
	}

	response, err := c.ExperienceUseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Error(err)
		return err
	}

	return ctx.JSON(model.WebResponse[model.ExperienceResponse]{
		Success: true,
		Message: "success create experience",
		Data:    *response,
	})
}

// Search Experience

func (c *ExperienceController) Search(ctx *fiber.Ctx) error {
	request := new(model.SearchExperienceRequest)
	request.Key = ctx.Query("key", "")
	request.Type = ctx.Query("type", "")
	request.Page = ctx.QueryInt("page", 1)
	request.Size = ctx.QueryInt("size", 10)

	responses, total, err := c.ExperienceUseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching contact")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.ExperienceResponse]{
		Data:   responses,
		Paging: paging,
	})
}
