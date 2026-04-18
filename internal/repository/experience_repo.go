package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"id.diengs.backend/internal/entity"
	"id.diengs.backend/internal/model"
)

type ExperienceRepo struct {
	Repository[entity.Experience]
	Log *logrus.Logger
}

func NewExperienceRepo(log *logrus.Logger) *ExperienceRepo {
	return &ExperienceRepo{
		Log: log,
	}
}

func (r *ExperienceRepo) FindByIdWithImages(db *gorm.DB, experience *entity.Experience, id string) error {
	return db.Preload("Images").Where("id = ?", id).Take(experience).Error
}

func (r *ExperienceRepo) Search(db *gorm.DB, request *model.SearchExperienceRequest) ([]entity.Experience, int64, error) {
	var experience []entity.Experience
	if err := db.Scopes(r.FilterSearch(request)).Offset((request.Page - 1) * request.Size).Limit(request.Size).Find(&experience).Error; err != nil {
		return nil, 0, err
	}

	var total int64 = 0
	if err := db.Model(&entity.Experience{}).Scopes(r.FilterSearch(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return experience, total, nil
}

func (r *ExperienceRepo) FilterSearch(request *model.SearchExperienceRequest) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if request.Type != "" {
			tx = tx.Where("experience_type = ?", request.Type)
		}

		if request.Key != "" {
			key := "%" + request.Key + "%"
			tx = tx.Where(
				"(title ILIKE ? OR address ILIKE ? OR description ILIKE ?)",
				key, key, key,
			)
		}

		return tx
	}
}
