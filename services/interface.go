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
	GetAllVisitations(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitations, int64, error)
	CreateVisitation(visitation models.Visitations, UserService IUserService) (models.Visitations, error)
	GetVisitationById(id int) (models.Visitations, error)
	UpdateVisitation(visitation models.Visitations, UserService IUserService) (models.Visitations, error)
	DeleteVisitation(id int) error
}

type IBorrowedService interface {
	GetAllBorroweds(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error)
	CreateBorrowed(borrowed models.Borrowed, UserService IUserService, BookService IBookService) (models.Borrowed, error)
	GetBorrowedById(id int) (models.Borrowed, error)
	UpdateBorrowed(borrowed models.Borrowed, UserService IUserService, BookService IBookService) (models.Borrowed, error)
	DeleteBorrowed(id int) error
}

type IBookReadsService interface {
	GetAllBookReads(query structs.Query, bookReadQuery structs.BookReadQuery) ([]models.BookRead, int64, error)
	CreateBookRead(bookRead models.BookRead) (models.BookRead, error)
	GetBookReadById(id int) (models.BookRead, error)
	UpdateBookRead(bookRead models.BookRead) (models.BookRead, error)
	DeleteBookRead(id int) error
}
