package repository

import (
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type VisitationRepository struct {
	db *gorm.DB
}

func NewVisitationRepository(DB *gorm.DB) IVisitationRepository {
	return &VisitationRepository{
		db: DB,
	}
}

func (v *VisitationRepository) GetById(id int) (models.Visitations, error) {
	var visitation models.Visitations
	err := v.db.Where("id = ?", id).First(&visitation).Error
	if err != nil {
		return models.Visitations{}, err
	}
	return visitation, nil
}

func (v *VisitationRepository) List(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitations, int64, error) {
	var visitations []models.Visitations
	var count int64
	startQuery := v.db
	offset := helpers.GetOffset(query)
	if visitationQuery.UserID != 0 {
		startQuery = startQuery.Where("user_id = ?", visitationQuery.UserID)
	}
	if visitationQuery.VisitedAtStart.Year() != 1 {
		startQuery = startQuery.Where("visited_at >= ?", visitationQuery.VisitedAtStart)
	}
	if visitationQuery.VisitedAtEnd.Year() != 1 {
		startQuery = startQuery.Where("visited_at <= ?", visitationQuery.VisitedAtEnd)
	}
	if visitationQuery.Duration != 0 {
		startQuery = startQuery.Where("duration = ?", visitationQuery.Duration)
	}
	err := startQuery.Limit(query.PerPage).Offset(offset).Find(&visitations).Error
	if err != nil {
		return []models.Visitations{}, count, err
	}
	err = v.db.Model(&visitations).Count(&count).Error
	if err != nil {
		return []models.Visitations{}, count, err
	}
	return visitations, count, nil
}

func (v *VisitationRepository) Create(visitation models.Visitations) (models.Visitations, error) {
	var Visitations models.Visitations = visitation
	err := v.db.Create(&Visitations).Error
	if err != nil {
		return models.Visitations{}, err
	}
	return Visitations, nil
}

func (v *VisitationRepository) Update(visitation models.Visitations) (models.Visitations, error) {
	var Visitations models.Visitations = visitation
	err := v.db.Model(&models.Visitations{}).Where("id = ?", Visitations.ID).Omit("id",  "created_at", "updated_at").Updates(&Visitations).Error
	if err != nil {
		return models.Visitations{}, err
	}
	return Visitations, nil
}

func (v *VisitationRepository) Delete(id int) error {
	return v.db.Delete(&models.Visitations{}, id).Error
}
