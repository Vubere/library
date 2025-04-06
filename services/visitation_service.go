package services

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	"victorubere/library/repository"
)

type VisitationService struct {
	repository repository.IVisitationRepository
}

func NewVisitationService(repository repository.IVisitationRepository) IVisitationService {
	return &VisitationService{
		repository: repository,
	}
}

func (v *VisitationService) GetAllVisitations(query structs.Query) ([]models.Visitation, error) {
	visitations, err := v.repository.List(query)
	if err != nil {
		return []models.Visitation{}, err
	}
	return visitations, nil
}

func (v *VisitationService) CreateVisitation(visitation models.Visitation) (models.Visitation, error) {
	createdVisitation, err := v.repository.Create(visitation)
	if err != nil {
		return models.Visitation{}, err
	}
	return createdVisitation, nil
}

func (v *VisitationService) GetVisitationById(id int) (models.Visitation, error) {
	visitation, err := v.repository.GetById(id)
	if err != nil {
		return models.Visitation{}, err
	}
	return visitation, nil
}

func (v *VisitationService) UpdateVisitation(visitation models.Visitation) (models.Visitation, error) {
	updatedVisitation, err := v.repository.Update(visitation)
	if err != nil {
		return models.Visitation{}, err
	}
	return updatedVisitation, nil
}

func (v *VisitationService) DeleteVisitation(id int) error {
	err := v.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
