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
	List(query structs.Query) ([]models.Visitation, error)
	Create(visitation models.Visitation) (models.Visitation, error)
	Update(visitation models.Visitation) (models.Visitation, error)
	Delete(id int) error
}

type IReservationRepository interface {
	GetById(id int) (models.Reservation, error)
	List(query structs.Query) ([]models.Reservation, error)
	Create(reservation models.Reservation) (models.Reservation, error)
	Update(reservation models.Reservation) (models.Reservation, error)
	Delete(id int) error
}
