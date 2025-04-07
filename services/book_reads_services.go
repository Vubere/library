package services

import (
	"errors"
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	repository "victorubere/library/repository"
)

type BookReadsService struct {
	repository repository.IBookReadssRepository
}

func NewBookReadsService(repository repository.IBookReadssRepository) IBookReadssService {
	return &BookReadsService{
		repository: repository,
	}
}

func (r *BookReadsService) GetAllBookReadss(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]models.BookReads, int64, error) {
	bookReads, count, err := r.repository.List(query, bookReadQuery)
	if err != nil {
		return []models.BookReads{}, 0, err
	}
	return bookReads, count, nil
}

func (r *BookReadsService) CreateBookReads(bookRead models.BookReads, userService IUserService, bookService IBookService, visitationService IVisitationService) (models.BookReads, error) {
	_, err := userService.GetUserById(int(bookRead.UserID))
	if err != nil {
		if err.Error() == "record not found" {
			return models.BookReads{}, errors.New("user not found")
		}
		return models.BookReads{}, err
	}
	_, err = bookService.GetBookById(int(bookRead.BookID))
	if err != nil {
		if err.Error() == "record not found" {
			return models.BookReads{}, errors.New("book not found")
		}
		return models.BookReads{}, err
	}
	_, err = visitationService.GetVisitationById(int(bookRead.VisitationID))
	if err != nil {
		if err.Error() == "record not found" {
			return models.BookReads{}, errors.New("visitation not found")
		}
		return models.BookReads{}, err
	}
	createdBookReads, err := r.repository.Create(bookRead)
	if err != nil {
		return models.BookReads{}, err
	}
	return createdBookReads, nil
}

func (r *BookReadsService) GetBookReadsById(id int) (models.BookReads, error) {
	bookRead, err := r.repository.GetById(id)
	if err != nil {
		return models.BookReads{}, err
	}
	return bookRead, nil
}

func (r *BookReadsService) UpdateBookReads(bookRead models.BookReads, userService IUserService, bookService IBookService, visitationService IVisitationService) (models.BookReads, error) {
	if bookRead.UserID != 0 {
		_, err := userService.GetUserById(int(bookRead.UserID))
		if err != nil {
			if err.Error() == "record not found" {
				return models.BookReads{}, errors.New("user not found")
			}
			return models.BookReads{}, err
		}
	}

	if bookRead.BookID != 0 {
		_, err := bookService.GetBookById(int(bookRead.BookID))
		if err != nil {
			if err.Error() == "record not found" {
				return models.BookReads{}, errors.New("book not found")
			}
			return models.BookReads{}, err
		}
	}

	if bookRead.VisitationID != 0 {
		_, err := visitationService.GetVisitationById(int(bookRead.VisitationID))
		if err != nil {
			if err.Error() == "record not found" {
				return models.BookReads{}, errors.New("visitation not found")
			}
			return models.BookReads{}, err
		}
	}
	updatedBookReads, err := r.repository.Update(bookRead)
	if err != nil {
		return models.BookReads{}, err
	}
	return updatedBookReads, nil
}

func (r *BookReadsService) DeleteBookReads(id int) error {
	err := r.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
