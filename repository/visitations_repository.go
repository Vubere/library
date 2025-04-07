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

func (v *VisitationRepository) GetById(id int) (models.Visitation, error) {
	var visitation models.Visitation
	err := v.db.Where("visitations.id = ?", id).Preload("User").First(&visitation).Error
	if err != nil {
		return models.Visitation{}, err
	}
	return visitation, nil
}

func (v *VisitationRepository) List(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitation, int64, error) {
	var visitations []models.Visitation
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
	startQuery = startQuery.Select("visitations.*, visitations.id as id").Joins("LEFT JOIN users ON users.id = visitations.user_id")
	err := startQuery.Preload("User").Limit(query.PerPage).Offset(offset).Find(&visitations).Error	
	if err != nil {
		return []models.Visitation{}, count, err
	}
	err = startQuery.Count(&count).Error
	if err != nil {
		return []models.Visitation{}, count, err
	}
	return visitations, count, nil
}

func (v *VisitationRepository) Create(visitation models.Visitation) (models.Visitation, error) {
	var Visitation models.Visitation = visitation
	err := v.db.Create(&Visitation).Error
	if err != nil {
		return models.Visitation{}, err
	}
	return Visitation, nil
}

func (v *VisitationRepository) Update(visitation models.Visitation) (models.Visitation, error) {
	var Visitation models.Visitation = visitation
	err := v.db.Model(&models.Visitation{}).Where("id = ?", Visitation.ID).Omit("id", "created_at", "updated_at").Updates(&Visitation).Error
	if err != nil {
		return models.Visitation{}, err
	}
	return Visitation, nil
}

func (v *VisitationRepository) Delete(id int) error {
	return v.db.Table("visitations").Where("id = ?", id).Delete(&models.Visitation{}).Error
}
