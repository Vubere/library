package repository

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
)

type IUserRepository interface {
	GetById(id int) (models.User, error)
	List(query structs.Query, userDetailsQuery structs.UserQuery) ([]models.User, int64, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id int) error
	GetByEmail(email string) (models.User, error)
}

type IBookRepository interface {
	GetById(id int) (models.Book, error)
	List(query structs.Query, bookQuery structs.BookQuery) ([]models.Book, int64, error)
	Create(book models.Book) (models.Book, error)
	Update(book models.Book) (models.Book, error)
	Delete(id int) error
}

type IVisitationRepository interface {
	GetById(id int) (models.Visitation, error)
	List(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitation, int64, error)
	Create(visitation models.Visitation) (models.Visitation, error)
	Update(visitation models.Visitation) (models.Visitation, error)
	Delete(id int) error
}

type IBorrowedRepository interface {
	GetById(id int) (models.Borrowed, error)
	List(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error)
	Create(borrowed models.Borrowed) (models.Borrowed, error)
	Update(borrowed models.Borrowed) (models.Borrowed, error)
	Delete(id int) error
}

type IBookReadssRepository interface {
	GetById(id int) (models.BookReads, error)
	List(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]models.BookReads, int64, error)
	Create(bookRead models.BookReads) (models.BookReads, error)
	Update(bookRead models.BookReads) (models.BookReads, error)
	Delete(id int) error
}
