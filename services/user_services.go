package services

import (
	"victorubere/library/lib/structs"
	"victorubere/library/models"
	repository "victorubere/library/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) RegisterUser(user models.User) (models.User, error) {
	return u.userRepository.Create(user)
	// TODO: Implement the RegisterUser method.
}

func (u *UserService) LoginUser(user models.User) (models.User, error) {
	return u.userRepository.GetByEmail(user.Email)
	// TODO: Implement the LoginUser method.
}

func (u *UserService) GetUserById(id int) (models.User, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserService) GetAllUsers(query structs.Query, userDetailsQuery structs.UserQuery) ([]models.User, int64, error) {
	users, count, err := u.userRepository.List(query, userDetailsQuery)
	if err != nil {
		return []models.User{}, 0, err
	}
	return users, count, nil
}

func (u *UserService) DeleteUser(id int) error {
	return u.userRepository.Delete(id)
	// TODO: Implement the DeleteUser method.
}

func (u *UserService) UpdateUser(user models.User) (models.User, error) {
	updatedUser, err := u.userRepository.Update(user)
	if err != nil {
		return models.User{}, err
	}
	return updatedUser, nil
}

func (u *UserService) GetUserByEmail(email string) (models.User, error) {
	return u.userRepository.GetByEmail(email)
}

func (u *UserService) GetTotalUsers(userDetailsQuery structs.UserQuery) (int64, error) {
	return u.userRepository.TotalUsers(userDetailsQuery)
}

func (u *UserService) GetUserSummary(id int, visitationService IVisitationService, borrowedService IBorrowedService, bookReadService IBookReadsService) (structs.UserSummaryDTO, error) {
	user, err := u.GetUserById(id)
	var returnedUser models.UserDTO
	returnedUser.ID = user.ID
	returnedUser.Email = user.Email
	returnedUser.PhoneNumber = user.PhoneNumber
	returnedUser.Gender = user.Email
	var hasVisited bool = false
	var hasBorrowed bool = false
	var hasRead bool = false
	if err != nil {
		return structs.UserSummaryDTO{}, err
	}
	visitationsCount, err := visitationService.GetTotalVisitations(structs.VisitationQuery{UserID: id})
	if err != nil {
		if err.Error() == "record not found" {
			hasVisited = false
		} else {
			return structs.UserSummaryDTO{}, err
		}
	}
	borrowedsCount, err := borrowedService.GetTotalBorrowings(structs.BorrowedQuery{UserID: id})
	if err != nil {
		if err.Error() == "record not found" {
			hasBorrowed = false
		} else {
			return structs.UserSummaryDTO{}, err
		}
	}
	bookReadsCount, err := bookReadService.GetTotalBookReads(structs.BookReadsQuery{UserID: id})
	if err != nil {
		if err.Error() == "record not found" {
			hasRead = false
		} else {
			return structs.UserSummaryDTO{}, err
		}
	}
	mostReadBook, err := bookReadService.GetMostReadBooks(structs.Query{Page: 1, PerPage: 1}, structs.BookReadsQuery{UserID: id})
	if err != nil {
		if err.Error() == "record not found" {
			hasRead = false
		} else {
			return structs.UserSummaryDTO{}, err
		}
	}
	mostBorrowedBook, err := borrowedService.GetMostBorrowedBooks(structs.BorrowedQuery{UserID: id})
	if err != nil {
		if err.Error() == "record not found" {
			hasBorrowed = false
		} else {
			return structs.UserSummaryDTO{}, err
		}
		return structs.UserSummaryDTO{}, err
	}
	var topMostReadBook *structs.MostBookReadsDTO
	if len(mostReadBook) > 0 {
		topMostReadBook = &mostReadBook[0]
	}
	var topMostBorrowedBook *structs.MostBorrowedBookDTO
	if mostBorrowedBook.BookBorrowedCount > 0 {
		topMostBorrowedBook = &mostBorrowedBook
	}
	if !hasVisited {
		visitationsCount = 0
	}
	if !hasBorrowed {
		borrowedsCount = 0
	}
	if !hasRead {
		bookReadsCount = 0
	}

	return structs.UserSummaryDTO{
		UserDetails:      returnedUser,
		VisitationsCount: visitationsCount,
		BorrowedsCount:   borrowedsCount,
		BookReadsCount:   bookReadsCount,
		MostReadBook:     topMostReadBook,
		MostBorrowedBook: topMostBorrowedBook,
	}, nil
}
