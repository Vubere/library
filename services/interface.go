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
	GetAllBooks(query structs.Query, bookQuery structs.BookQuery) ([]models.Book,int64, error)
	CreateBook(book models.Book) (models.Book, error)
	GetBookById(id int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(id int) error
}

type IVisitationService interface {
	GetAllVisitations(query structs.Query) ([]models.Visitation, error)
	CreateVisitation(visitation models.Visitation) (models.Visitation, error)
	GetVisitationById(id int) (models.Visitation, error)
	UpdateVisitation(visitation models.Visitation) (models.Visitation, error)
	DeleteVisitation(id int) error
}

type IReservationService interface {
	GetAllReservations(query structs.Query) ([]models.Reservation, error)
	CreateReservation(reservation models.Reservation) (models.Reservation, error)
	GetReservationById(id int) (models.Reservation, error)
	UpdateReservation(reservation models.Reservation) (models.Reservation, error)
	DeleteReservation(id int) error
}
