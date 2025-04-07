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
}

type IBookService interface {
	GetAllBooks(query structs.Query, bookQuery structs.BookQuery) ([]models.Book, int64, error)
	CreateBook(book models.Book) (models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(id int) error
}

type IVisitationService interface {
	GetAllVisitation(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitation, int64, error)
	CreateVisitation(visitation models.Visitation, UserService IUserService) (models.Visitation, error)
	GetVisitationById(id int) (models.Visitation, error)
	UpdateVisitation(visitation models.Visitation, UserService IUserService) (models.Visitation, error)
	DeleteVisitation(id int) error
}

type IBorrowedService interface {
	GetAllBorroweds(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error)
	CreateBorrowed(borrowed models.Borrowed, UserService IUserService, BookService IBookService) (models.Borrowed, error)
	GetBorrowedById(id int) (models.Borrowed, error)
	UpdateBorrowed(borrowed models.Borrowed, UserService IUserService, BookService IBookService) (models.Borrowed, error)
	DeleteBorrowed(id int) error
}

type IBookReadssService interface {
	GetAllBookReadss(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]models.BookReads, int64, error)
	CreateBookReads(bookRead models.BookReads, UserService IUserService, BookService IBookService, VisitationService IVisitationService) (models.BookReads, error)
	GetBookReadsById(id int) (models.BookReads, error)
	UpdateBookReads(bookRead models.BookReads, UserService IUserService, BookService IBookService, VisitationService IVisitationService) (models.BookReads, error)
	DeleteBookReads(id int) error
}
