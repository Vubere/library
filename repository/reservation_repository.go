package repository

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type BorrowedRepository struct {
	db *gorm.DB
}

func NewBorrowedRepository(DB *gorm.DB) IBorrowedRepository {
	return &BorrowedRepository{
		db: DB,
	}
}

func (r *BorrowedRepository) GetById(id int) (models.Borrowed, error) {
	return models.Borrowed{}, nil
}

func (r *BorrowedRepository) List(query structs.Query) ([]models.Borrowed, error) {
	return []models.Borrowed{}, nil
}

func (r *BorrowedRepository) Create(borrowed models.Borrowed) (models.Borrowed, error) {
	return models.Borrowed{}, nil
}

func (r *BorrowedRepository) Update(borrowed models.Borrowed) (models.Borrowed, error) {
	return models.Borrowed{}, nil
}

func (r *BorrowedRepository) Delete(id int) error {
	return nil
}
