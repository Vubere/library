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
