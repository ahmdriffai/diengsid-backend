package model

import "id.diengs.backend/internal/entity"

// Request

type ExperienceCreateImageRequest struct {
	ImageURL  string `json:"image_url" validate:"required"`
	IsPrimary bool   `json:"is_primary"`
}

type ExperienceCreateRequest struct {
	ExperienceType string                         `json:"experience_type" validate:"required"`
	Title          string                         `json:"title" validate:"required"`
	Address        string                         `json:"address" validate:"required"`
	Description    string                         `json:"description" validate:"required"`
	ThumbnailURL   *string                        `json:"thumbnail_url"`
	Lat            *float64                       `json:"lat"`
	Lng            *float64                       `json:"lng"`
	BasePrice      float64                        `json:"base_price" validate:"required"`
	Images         []ExperienceCreateImageRequest `json:"images"`
}

type SearchExperienceRequest struct {
	Key  string `json:"key" validate:"max=100"`
	Type string `json:"type"`
	Page int    `json:"page" validate:"min=1"`
	Size int    `json:"size" validate:"min=1,max=100"`
}

// Response

type ExperienceImageResponse struct {
	ID           string `json:"id,omitempty"`
	ExperienceID string `json:"experience_id,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	IsPrimary    bool   `json:"is_primary,omitempty"`
	CreatedAt    int64  `json:"created_at,omitempty"`
	UpdatedAt    int64  `json:"updated_at,omitempty"`
}

type ExperienceResponse struct {
	ID             string                    `json:"id,omitempty"`
	ExperienceType string                    `json:"experience_type,omitempty"`
	Title          string                    `json:"title,omitempty"`
	Address        string                    `json:"address,omitempty"`
	Description    string                    `json:"description,omitempty"`
	ThumbnailURL   *string                   `json:"thumbnail_url,omitempty"`
	Lat            *float64                  `json:"lat,omitempty"`
	Lng            *float64                  `json:"lng,omitempty"`
	BasePrice      float64                   `json:"base_price,omitempty"`
	Images         []ExperienceImageResponse `json:"images,omitempty"`
	CreatedAt      int64                     `json:"created_at,omitempty"`
	UpdatedAt      int64                     `json:"updated_at,omitempty"`
}

func ExperienceToResponse(experience *entity.Experience) *ExperienceResponse {
	if experience == nil {
		return nil
	}

	images := make([]ExperienceImageResponse, 0, len(experience.Images))
	for _, image := range experience.Images {
		images = append(images, ExperienceImageResponse{
			ID:           image.ID,
			ExperienceID: image.ExperienceID,
			ImageURL:     image.ImageURL,
			IsPrimary:    image.IsPrimary,
			CreatedAt:    image.CreatedAt,
			UpdatedAt:    image.UpdatedAt,
		})
	}

	return &ExperienceResponse{
		ID:             experience.ID,
		ExperienceType: experience.ExperienceType,
		Title:          experience.Title,
		Address:        experience.Address,
		Description:    experience.Description,
		ThumbnailURL:   experience.ThumbnailURL,
		Lat:            experience.Lat,
		Lng:            experience.Lng,
		BasePrice:      experience.BasePrice,
		Images:         images,
		CreatedAt:      experience.CreatedAt,
		UpdatedAt:      experience.UpdatedAt,
	}
}
