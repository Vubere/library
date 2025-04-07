package services

import (
	"errors"
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	repository "victorubere/library/repository"
)

type BookReadService struct {
	repository repository.IBookReadsRepository
}

func NewBookReadService(repository repository.IBookReadsRepository) IBookReadsService {
	return &BookReadService{
		repository: repository,
	}
}

func (r *BookReadService) GetAllBookReads(query structs.Query, bookReadQuery structs.BookReadQuery) ([]models.BookRead, int64, error) {
	bookReads, count, err := r.repository.List(query, bookReadQuery)
	if err != nil {
		return []models.BookRead{}, 0, err
	}
	return bookReads, count, nil
}

func (r *BookReadService) CreateBookRead(bookRead models.BookRead, userService IUserService, bookService IBookService, visitationService IVisitationService) (models.BookRead, error) {
	_, err := userService.GetUserById(int(bookRead.UserID))
	if err != nil {
		if err.Error() == "record not found" {
			return models.BookRead{}, errors.New("user not found")
		}
		return models.BookRead{}, err
	}
	_, err = bookService.GetBookById(int(bookRead.BookID))
	if err != nil {
		if err.Error() == "record not found" {
			return models.BookRead{}, errors.New("book not found")
		}
		return models.BookRead{}, err
	}
	_, err = visitationService.GetVisitationById(int(bookRead.VisitationID))
	if err != nil {
		if err.Error() == "record not found" {
			return models.BookRead{}, errors.New("visitation not found")
		}
		return models.BookRead{}, err
	}
	createdBookRead, err := r.repository.Create(bookRead)
	if err != nil {
		return models.BookRead{}, err
	}
	return createdBookRead, nil
}

func (r *BookReadService) GetBookReadById(id int) (models.BookRead, error) {
	bookRead, err := r.repository.GetById(id)
	if err != nil {
		return models.BookRead{}, err
	}
	return bookRead, nil
}

func (r *BookReadService) UpdateBookRead(bookRead models.BookRead, userService IUserService, bookService IBookService, visitationService IVisitationService) (models.BookRead, error) {
	if bookRead.UserID != 0 {
		_, err := userService.GetUserById(int(bookRead.UserID))
		if err != nil {
			if err.Error() == "record not found" {
				return models.BookRead{}, errors.New("user not found")
			}
			return models.BookRead{}, err
		}
	}

	if bookRead.BookID != 0 {
		_, err := bookService.GetBookById(int(bookRead.BookID))
		if err != nil {
			if err.Error() == "record not found" {
				return models.BookRead{}, errors.New("book not found")
			}
			return models.BookRead{}, err
		}
	}

	if bookRead.VisitationID != 0 {
		_, err := visitationService.GetVisitationById(int(bookRead.VisitationID))
		if err != nil {
			if err.Error() == "record not found" {
				return models.BookRead{}, errors.New("visitation not found")
			}
			return models.BookRead{}, err
		}
	}
	updatedBookRead, err := r.repository.Update(bookRead)
	if err != nil {
		return models.BookRead{}, err
	}
	return updatedBookRead, nil
}

func (r *BookReadService) DeleteBookRead(id int) error {
	err := r.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
