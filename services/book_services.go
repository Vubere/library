package services

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	repository "victorubere/library/repository"
)

type BookService struct {
	bookRepository repository.IBookRepository
}

func NewBookService(bookRepository repository.IBookRepository) IBookService {
	return &BookService{
		bookRepository: bookRepository,
	}
}

func (b *BookService) GetAllBooks(query structs.Query, bookQuery structs.BookQuery) ([]models.Book, int64, error) {
	return b.bookRepository.List(query, bookQuery)
}

func (b *BookService) CreateBook(book models.Book) (models.Book, error) {
	return b.bookRepository.Create(book)
}

func (b *BookService) GetBookById(id int) (models.Book, error) {
	return b.bookRepository.GetById(id)
}

func (b *BookService) UpdateBook(book models.Book) (models.Book, error) {
	return b.bookRepository.Update(book)
}

func (b *BookService) DeleteBook(id int) error {
	return b.bookRepository.Delete(id)
}
