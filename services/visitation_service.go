package services

import (
	"errors"
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

func (v *VisitationService) GetAllVisitations(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitations, int64, error) {
	visitations, count, err := v.repository.List(query, visitationQuery)
	if err != nil {
		return []models.Visitations{}, 0, err
	}
	return visitations, count, nil
}

func (v *VisitationService) CreateVisitation(visitation models.Visitations, UserService IUserService) (models.Visitations, error) {
	_, err := UserService.GetUserById(visitation.UserId)
	if err != nil {
		if err.Error() == "record not found" {
			return models.Visitations{}, errors.New("user not found")
		}
		return models.Visitations{}, err
	}
	createdVisitation, err := v.repository.Create(visitation)
	if err != nil {
		return models.Visitations{}, err
	}
	return createdVisitation, nil
}

func (v *VisitationService) GetVisitationById(id int) (models.Visitations, error) {
	visitation, err := v.repository.GetById(id)
	if err != nil {
		return models.Visitations{}, err
	}
	return visitation, nil
}

func (v *VisitationService) UpdateVisitation(visitation models.Visitations, UserService IUserService) (models.Visitations, error) {
	if visitation.UserId != 0 {
		_, err := UserService.GetUserById(visitation.UserId)
		if err != nil {
			if err.Error() == "record not found" {
				return models.Visitations{}, errors.New("user not found")
			}
			return models.Visitations{}, err
		}
	}
	updatedVisitation, err := v.repository.Update(visitation)
	if err != nil {
		return models.Visitations{}, err
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
