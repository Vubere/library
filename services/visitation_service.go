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

func (v *VisitationService) GetAllVisitation(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitation, int64, error) {
	visitations, count, err := v.repository.List(query, visitationQuery)
	if err != nil {
		return []models.Visitation{}, 0, err
	}
	return visitations, count, nil
}

func (v *VisitationService) CreateVisitation(visitation models.Visitation, UserService IUserService) (models.Visitation, error) {
	_, err := UserService.GetUserById(visitation.UserId)
	if err != nil {
		if err.Error() == "record not found" {
			return models.Visitation{}, errors.New("user not found")
		}
		return models.Visitation{}, err
	}
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

func (v *VisitationService) UpdateVisitation(visitation models.Visitation, UserService IUserService) (models.Visitation, error) {
	if visitation.UserId != 0 {
		_, err := UserService.GetUserById(visitation.UserId)
		if err != nil {
			if err.Error() == "record not found" {
				return models.Visitation{}, errors.New("user not found")
			}
			return models.Visitation{}, err
		}
	}
	updatedVisitation, err := v.repository.Update(visitation)
	if err != nil {
		return models.Visitation{}, err
	}
	return updatedVisitation, nil
}

func (v *VisitationService) DeleteVisitation(id int) error {
	if _, err := v.GetVisitationById(id); err != nil {
		return err
	}
	err := v.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (v *VisitationService) GetTotalVisitations(visitationQuery structs.VisitationQuery) (int64, error) {
	totalVisitations, err := v.repository.TotalVisitations(visitationQuery)
	if err != nil {
		return 0, err
	}
	return totalVisitations, nil
}
