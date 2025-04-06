package services

import (
	"errors"
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

func (r *BorrowedService) GetAllBorroweds(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error) {
	borroweds, count, err := r.repository.List(query, borrowedQuery)
	if err != nil {
		return []models.Borrowed{}, 0, err
	}
	return borroweds, count, nil
}

func (r *BorrowedService) CreateBorrowed(borrowed models.Borrowed, userService IUserService, bookService IBookService) (models.Borrowed, error) {
	_, err := userService.GetUserById(borrowed.UserId)
	if err != nil {
		if err.Error() == "record not found" {
			return models.Borrowed{}, errors.New("user not found")
		}
		return models.Borrowed{}, err
	}
	_, err = bookService.GetBookById(borrowed.BookId)
	if err != nil {
		if err.Error() == "record not found" {
			return models.Borrowed{}, errors.New("book not found")
		}
		return models.Borrowed{}, err
	}
	createdBorrowed, err := r.repository.Create(borrowed)
	if err != nil {
		return models.Borrowed{}, err
	}
	return createdBorrowed, nil
}

func (r *BorrowedService) GetBorrowedById(id int) (models.Borrowed, error) {	
	return r.repository.GetById(id)
}

func (r *BorrowedService) UpdateBorrowed(borrowed models.Borrowed, userService IUserService, bookService IBookService) (models.Borrowed, error) {
	if borrowed.UserId != 0 {
		_, err := userService.GetUserById(borrowed.UserId)
		if err != nil {
			if err.Error() == "record not found" {
				return models.Borrowed{}, errors.New("user not found")
			}
			return models.Borrowed{}, err
		}
	}
	
	if borrowed.BookId != 0 {
		_, err := bookService.GetBookById(borrowed.BookId)
		if err != nil {
			if err.Error() == "record not found" {
				return models.Borrowed{}, errors.New("book not found")
			}
			return models.Borrowed{}, err
		}
	}
	updatedBorrowed, err := r.repository.Update(borrowed)
	if err != nil {
		return models.Borrowed{}, err
	}
	return updatedBorrowed, nil
}

func (r *BorrowedService) DeleteBorrowed(id int) error {
	_, err := r.repository.GetById(id)
	if err != nil {
		return err
	}
	return r.repository.Delete(id)
}
