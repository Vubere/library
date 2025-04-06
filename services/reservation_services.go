package services

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	"victorubere/library/repository"
)

type ReservationService struct {
	repository repository.IReservationRepository
}

func NewReservationService(repository repository.IReservationRepository) IReservationService {
	return &ReservationService{
		repository: repository,
	}
}

func (r *ReservationService) GetAllReservations(query structs.Query) ([]models.Reservation, error) {
	// TODO: Implement the GetAllReservations method.
	return []models.Reservation{}, nil
}

func (r *ReservationService) CreateReservation(reservation models.Reservation) (models.Reservation, error) {
	// TODO: Implement the CreateReservation method.
	return models.Reservation{}, nil
}

func (r *ReservationService) GetReservationById(id int) (models.Reservation, error) {
	// TODO: Implement the GetReservationById method.
	return models.Reservation{}, nil
}

func (r *ReservationService) UpdateReservation(reservation models.Reservation) (models.Reservation, error) {
	// TODO: Implement the UpdateReservation method.
	return models.Reservation{}, nil
}

func (r *ReservationService) DeleteReservation(id int) error {
	// TODO: Implement the DeleteReservation method.
	return nil
}
