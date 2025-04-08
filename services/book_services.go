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

func (b *BookService) GetBookSummaryDTO(id int, visitationService IVisitationService, borrowedService IBorrowedService, bookReadService IBookReadsService) (structs.BookSummaryDTO, error) {
	book, err := b.GetBookById(id)
	if err != nil {
		return structs.BookSummaryDTO{}, err
	}
	readsCount, err := bookReadService.GetTotalBookReads(structs.BookReadsQuery{BookID: id})
	if err != nil {
		return structs.BookSummaryDTO{}, err
	}
	borrowedsCount, err := borrowedService.GetTotalBorrowings(structs.BorrowedQuery{BookID: id})
	if err != nil {
		return structs.BookSummaryDTO{}, err
	}
	readMostByUser, err := bookReadService.GetUserWithMostBookReads(int(book.ID), structs.BookReadsQuery{BookID: id})
	if err != nil {
		return structs.BookSummaryDTO{}, err
	}
	borrowedMostByUser, err := borrowedService.GetUserWhoBorrowedBookMost(int(book.ID), structs.BorrowedQuery{BookID: id})
	if err != nil {
		return structs.BookSummaryDTO{}, err
	}

	return structs.BookSummaryDTO{
		BookDetails:           book,
		ReadsCount:            readsCount,
		BorrowedsCount:        borrowedsCount,
		ReadMostByUser:        readMostByUser,
		BorrowedMostByUserDTO: borrowedMostByUser,
	}, nil
}

func (b *BookService) GetBooksSummaryDTO(query structs.Query, bookQuery structs.BookQuery, visitationService IVisitationService, borrowedService IBorrowedService, bookReadService IBookReadsService) (structs.BooksSummaryDTO, error) {
	var bookSummary structs.BooksSummaryDTO
	var err error
	bookSummary.BooksCount, err = b.GetTotalBooks(structs.BookQuery{})
	if err != nil {
		return structs.BooksSummaryDTO{}, err
	}
	mostRead, err := bookReadService.GetMostReadBooks(structs.Query{Page: 1, PerPage: 1}, structs.BookReadsQuery{})
	if err != nil {
		if err.Error() == "record not found" {
			bookSummary.MostReadBook = structs.MostBookReadsDTO{}
		} else {
			return structs.BooksSummaryDTO{}, err
		}
	}
	bookSummary.MostReadBook = mostRead[0]
	mostBorrowed, err := borrowedService.GetMostBorrowedBooks(structs.BorrowedQuery{})
	if err != nil {
		if err.Error() == "record not found" {
			bookSummary.MostBorrowedBookDTO = structs.MostBorrowedBookDTO{}
		} else {
			return structs.BooksSummaryDTO{}, err
		}
	}
	bookSummary.MostBorrowedBookDTO = mostBorrowed
	return bookSummary, nil
}

func (b *BookService) GetTotalBooks(bookQuery structs.BookQuery) (int64, error) {
	return b.bookRepository.TotalBooks(bookQuery)
}
