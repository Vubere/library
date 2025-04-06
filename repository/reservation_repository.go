package repository

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(DB *gorm.DB) IReservationRepository {
	return &ReservationRepository{
		db: DB,
	}
}

func (r *ReservationRepository) GetById(id int) (models.Reservation, error) {
	return models.Reservation{}, nil
}

func (r *ReservationRepository) List(query structs.Query) ([]models.Reservation, error) {
	return []models.Reservation{}, nil
}

func (r *ReservationRepository) Create(reservation models.Reservation) (models.Reservation, error) {
	return models.Reservation{}, nil
}

func (r *ReservationRepository) Update(reservation models.Reservation) (models.Reservation, error) {
	return models.Reservation{}, nil
}

func (r *ReservationRepository) Delete(id int) error {
	return nil
}
