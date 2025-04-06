package repository

import (
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
	return models.Visitation{}, nil
}

func (v *VisitationRepository) List(query structs.Query) ([]models.Visitation, error) {
	return []models.Visitation{}, nil
}

func (v *VisitationRepository) Create(visitation models.Visitation) (models.Visitation, error) {
	return models.Visitation{}, nil
}

func (v *VisitationRepository) Update(visitation models.Visitation) (models.Visitation, error) {
	return models.Visitation{}, nil
}

func (v *VisitationRepository) Delete(id int) error {
	return nil
}
