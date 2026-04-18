package usecase

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"id.diengs.backend/internal/entity"
	"id.diengs.backend/internal/model"
	"id.diengs.backend/internal/repository"
)

type ExperienceUseCase struct {
	DB                  *gorm.DB
	Log                 *logrus.Logger
	Validate            *validator.Validate
	ExperienceRepo      *repository.ExperienceRepo
	ExperienceImageRepo *repository.ExperienceImageRepo
}

func NewExperienceUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	validate *validator.Validate,
	experienceRepo *repository.ExperienceRepo,
	experienceImageRepo *repository.ExperienceImageRepo,
) *ExperienceUseCase {
	return &ExperienceUseCase{
		DB:                  db,
		Log:                 log,
		Validate:            validate,
		ExperienceRepo:      experienceRepo,
		ExperienceImageRepo: experienceImageRepo,
	}
}

// Create Experince
func (u *ExperienceUseCase) Create(ctx context.Context, request *model.ExperienceCreateRequest) (*model.ExperienceResponse, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("FAILED TO VALIDATE REQUEST.")
		return nil, fiber.ErrBadRequest
	}

	experience := &entity.Experience{
		ExperienceType: request.ExperienceType,
		Title:          request.Title,
		Address:        request.Address,
		Description:    request.Description,
		ThumbnailURL:   request.ThumbnailURL,
		Lat:            request.Lat,
		Lng:            request.Lng,
		BasePrice:      request.BasePrice,
	}

	if err := u.ExperienceRepo.Create(tx, experience); err != nil {
		u.Log.WithError(err).Error("FAILED TO CREATE EXPERIENCE.")
		return nil, fiber.ErrInternalServerError
	}

	for _, image := range request.Images {
		experienceImage := &entity.ExperienceImage{
			ExperienceID: experience.ID,
			ImageURL:     image.ImageURL,
			IsPrimary:    image.IsPrimary,
		}

		if err := u.ExperienceImageRepo.Create(tx, experienceImage); err != nil {
			u.Log.WithError(err).Error("FAILED TO CREATE EXPERIENCE IMAGE.")
			return nil, fiber.ErrInternalServerError
		}

		experience.Images = append(experience.Images, *experienceImage)
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.WithError(err).Error("FAILED TO COMMIT TRANSACTION.")
		return nil, fiber.ErrInternalServerError
	}

	return model.ExperienceToResponse(experience), nil
}

// Search Employee
func (c *ExperienceUseCase) Search(ctx context.Context, request *model.SearchExperienceRequest) ([]model.ExperienceResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	experiences, total, err := c.ExperienceRepo.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting experience")
		return nil, 0, fiber.ErrInternalServerError
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("Failed to commit transaction")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.ExperienceResponse, len(experiences))
	for i, experience := range experiences {
		responses[i] = *model.ExperienceToResponse(&experience)
	}

	return responses, total, nil
}
