package repository

import (
	"time"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRpository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetById(id int) (models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserRepository) List(query structs.Query, userDetailsQuery structs.UserQuery) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	dbExec := u.db
	offset := helpers.GetOffset(query)

	if userDetailsQuery.Email != "" {
		dbExec = dbExec.Where("email LIKE ?", "%"+userDetailsQuery.Email+"%")
	}
	if userDetailsQuery.Name != "" {
		dbExec = dbExec.Where("name LIKE ?", "%"+userDetailsQuery.Name+"%")
	}
	if userDetailsQuery.Gender != "" {
		dbExec = dbExec.Where("gender = ?", userDetailsQuery.Gender)
	}
	if userDetailsQuery.DateCreatedStart != "" {
		dateCreatedStart, _ := time.Parse("2006-01-02", userDetailsQuery.DateCreatedStart)
		dbExec = dbExec.Where("created_at >= ?", dateCreatedStart)
	}
	if userDetailsQuery.DateCreatedEnd != "" {
		dateCreatedEnd, _ := time.Parse("2006-01-02", userDetailsQuery.DateCreatedEnd)
		dbExec = dbExec.Where("created_at <= ?", dateCreatedEnd)
	}
	if userDetailsQuery.MinAge != 0 {
		age := time.Now().AddDate(-userDetailsQuery.MinAge, 0, 0)
		dbExec = dbExec.Where("date_of_birth <= ?", age)
	}

	if query.SortBy != "" && query.SortDirection != "" {
		dbExec = dbExec.Order(query.SortBy + " " + query.SortDirection)
	}
	err := dbExec.Limit(query.PerPage).Offset(offset).Find(&users).Error
	if err != nil {
		return []models.User{}, 0, err
	}
	err = dbExec.Model(&users).Count(&count).Error
	if err != nil {
		return []models.User{}, 0, err
	}
	return users, count, nil
}

func (u *UserRepository) Create(user models.User) (models.User, error) {

	err := u.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Update(user models.User) (models.User, error) {
	var User models.User = user
	err := u.db.Model(&models.User{}).Where("id = ?", User.ID).Omit("id", "created_at", "updated_at").Updates(&User).Error
	if err != nil {
		return models.User{}, err
	}
	return User, nil
}

func (u *UserRepository) Delete(id int) error {
	return u.db.Delete(&models.User{}, id).Error
}

func (u *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserRepository) TotalUsers(userDetailsQuery structs.UserQuery) (int64, error) {
	var count int64
	startQuery := u.db.Model(&models.User{})
	if userDetailsQuery.Email != "" {
		startQuery = startQuery.Where("email LIKE ?", "%"+userDetailsQuery.Email+"%")
	}
	if userDetailsQuery.Name != "" {
		startQuery = startQuery.Where("name LIKE ?", "%"+userDetailsQuery.Name+"%")
	}
	if userDetailsQuery.Gender != "" {
		startQuery = startQuery.Where("gender = ?", userDetailsQuery.Gender)
	}
	if userDetailsQuery.DateCreatedStart != "" {
		dateCreatedStart, _ := time.Parse("2006-01-02", userDetailsQuery.DateCreatedStart)
		startQuery = startQuery.Where("created_at >= ?", dateCreatedStart)
	}
	if userDetailsQuery.DateCreatedEnd != "" {
		dateCreatedEnd, _ := time.Parse("2006-01-02", userDetailsQuery.DateCreatedEnd)
		startQuery = startQuery.Where("created_at <= ?", dateCreatedEnd)
	}
	if userDetailsQuery.MinAge != 0 {
		age := time.Now().AddDate(-userDetailsQuery.MinAge, 0, 0)
		startQuery = startQuery.Where("date_of_birth <= ?", age)
	}
	err := startQuery.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
