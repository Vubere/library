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
	GetById(id int) (models.Visitations, error)
	List(query structs.Query, visitationQuery structs.VisitationQuery) ([]models.Visitations, int64,	 error)
	Create(visitation models.Visitations) (models.Visitations, error)
	Update(visitation models.Visitations) (models.Visitations, error)
	Delete(id int) error
}

type IBorrowedRepository interface {
	GetById(id int) (models.Borrowed, error)
	List(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error)
	Create(borrowed models.Borrowed) (models.Borrowed, error)
	Update(borrowed models.Borrowed) (models.Borrowed, error)
	Delete(id int) error
}

type IBookReadsRepository interface {
	GetById(id int) (models.BookRead, error)
	List(query structs.Query, bookReadQuery structs.BookReadQuery) ([]models.BookRead, int64, int64,	 error)
	Create(bookRead models.BookRead) (models.BookRead, error)
	Update(bookRead models.BookRead) (models.BookRead, error)
	Delete(id int) error
}
