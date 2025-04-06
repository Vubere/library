package services

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	"victorubere/library/repository"
)

type BorrowedService struct {
	repository repository.IBorrowedRepository
}

func NewBorrowedService(repository repository.IBorrowedRepository) IBorrowedService {
	return &BorrowedService{
		repository: repository,
	}
}

func (r *BorrowedService) GetAllBorroweds(query structs.Query) ([]models.Borrowed, error) {
	// TODO: Implement the GetAllBorroweds method.
	return []models.Borrowed{}, nil
}

func (r *BorrowedService) CreateBorrowed(borrowed models.Borrowed) (models.Borrowed, error) {
	// TODO: Implement the CreateBorrowed method.
	return models.Borrowed{}, nil
}

func (r *BorrowedService) GetBorrowedById(id int) (models.Borrowed, error) {
	// TODO: Implement the GetBorrowedById method.
	return models.Borrowed{}, nil
}

func (r *BorrowedService) UpdateBorrowed(borrowed models.Borrowed) (models.Borrowed, error) {
	// TODO: Implement the UpdateBorrowed method.
	return models.Borrowed{}, nil
}

func (r *BorrowedService) DeleteBorrowed(id int) error {
	// TODO: Implement the DeleteBorrowed method.
	return nil
}
