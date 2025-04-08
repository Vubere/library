package services

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
)

type IUserService interface {
	RegisterUser(User models.User) (models.User, error)
	LoginUser(User models.User) (models.User, error)
	GetUserById(id int) (models.User, error)
	GetAllUsers(query structs.Query, userDetailsQuery structs.UserQuery) ([]models.User, int64, error)
	DeleteUser(id int) error
	UpdateUser(User models.User) (models.User, error)
	GetUserSummary(id int, visitationService IVisitationService, borrowedService IBorrowedService, bookReadService IBookReadsService) (structs.UserSummaryDTO, error)
	GetUserByEmail(email string) (models.User, error)
	GetTotalUsers(userDetailsQuery structs.UserQuery) (int64, error)
}

type IBookService interface {
	GetAllBooks(query structs.Query, bookQuery structs.BookQuery) ([]models.Book, int64, error)
	CreateBook(book models.Book) (models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(id int) error
	GetBookSummaryDTO(id int, visitationService IVisitationService, borrowedService IBorrowedService, bookReadService IBookReadsService) (structs.BookSummaryDTO, error)
	GetBooksSummaryDTO(query structs.Query, bookQuery structs.BookQuery, visitationService IVisitationService, borrowedService IBorrowedService, bookReadService IBookReadsService) (structs.BooksSummaryDTO, error)
	GetTotalBooks(bookQuery structs.BookQuery) (int64, error)
}

type IVisitationService interface {
	GetAllVisitation(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitation, int64, error)
	CreateVisitation(visitation models.Visitation, UserService IUserService) (models.Visitation, error)
	GetVisitationById(id int) (models.Visitation, error)
	UpdateVisitation(visitation models.Visitation, UserService IUserService) (models.Visitation, error)
	DeleteVisitation(id int) error
	GetTotalVisitations(visitationQuery structs.VisitationQuery) (int64, error)
}

type IBorrowedService interface {
	GetAllBorroweds(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error)
	CreateBorrowed(borrowed models.Borrowed, UserService IUserService, BookService IBookService) (models.Borrowed, error)
	GetBorrowedById(id int) (models.Borrowed, error)
	UpdateBorrowed(borrowed models.Borrowed, UserService IUserService, BookService IBookService) (models.Borrowed, error)
	DeleteBorrowed(id int) error
	GetTotalBorrowings(borrowedQuery structs.BorrowedQuery) (int64, error)
	GetMostBorrowedBooks(borrowedQuery structs.BorrowedQuery) (structs.MostBorrowedBookDTO, error)
	GetUserWhoBorrowedBookMost(bookId int, borrowedQuery structs.BorrowedQuery) (structs.BorrowedMostByUserDTO, error)
}

type IBookReadsService interface {
	GetAllBookReadss(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]models.BookReads, int64, error)
	CreateBookReads(bookRead models.BookReads, UserService IUserService, BookService IBookService, VisitationService IVisitationService) (models.BookReads, error)
	GetBookReadsById(id int) (models.BookReads, error)
	UpdateBookReads(bookRead models.BookReads, UserService IUserService, BookService IBookService, VisitationService IVisitationService) (models.BookReads, error)
	DeleteBookReads(id int) error
	GetTotalBookReads(bookReadQuery structs.BookReadsQuery) (int64, error)
	GetMostReadBooks(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]structs.MostBookReadsDTO, error)
	GetUserWithMostBookReads(userId int, bookReadQuery structs.BookReadsQuery) (structs.BookReadMostByUserDTO, error)
}
