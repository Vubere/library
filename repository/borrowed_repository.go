package repository

import (
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type BorrowedRepository struct {
	db *gorm.DB
}

func NewBorrowedRepository(DB *gorm.DB) IBorrowedRepository {
	return &BorrowedRepository{
		db: DB,
	}
}

func (r *BorrowedRepository) GetById(id int) (models.Borrowed, error) {
	var borrowed models.Borrowed
	err := r.db.Where("id = ?", id).Select("borroweds.*, users.*, books.*, borroweds.id as id").Joins("LEFT JOIN users ON users.id = borroweds.user_id").Joins("LEFT JOIN books ON books.id = borroweds.book_id").First(&borrowed).Preload("User").Preload("Book").Error
	if err != nil {
		return models.Borrowed{}, err
	}
	return borrowed, nil
}

func (r *BorrowedRepository) List(query structs.Query, borrowedQuery structs.BorrowedQuery) ([]models.Borrowed, int64, error) {
	var borroweds []models.Borrowed
	var count int64
	startQuery := r.db
	offset := helpers.GetOffset(query)
	if borrowedQuery.UserID != 0 {
		startQuery = startQuery.Where("user_id = ?", borrowedQuery.UserID)
	}
	if borrowedQuery.BookID != 0 {
		startQuery = startQuery.Where("book_id = ?", borrowedQuery.BookID)
	}
	if borrowedQuery.Duration != 0 {
		startQuery = startQuery.Where("duration = ?", borrowedQuery.Duration)
	}
	if borrowedQuery.BorrowedAtStart.Year() != 1 {
		startQuery = startQuery.Where("borrowed_at >= ?", borrowedQuery.BorrowedAtStart)
	}
	if borrowedQuery.BorrowedAtEnd.Year() != 1 {
		startQuery = startQuery.Where("borrowed_at <= ?", borrowedQuery.BorrowedAtEnd)
	}
	if borrowedQuery.ReturnedAtStart.Year() != 1 {
		startQuery = startQuery.Where("returned_at >= ?", borrowedQuery.ReturnedAtStart)
	}
	if borrowedQuery.ReturnedAtEnd.Year() != 1 {
		startQuery = startQuery.Where("returned_at <= ?", borrowedQuery.ReturnedAtEnd)
	}
	startQuery = startQuery.Select("borroweds.*, borroweds.id as id").Joins("LEFT JOIN users ON users.id = borroweds.user_id").Joins("LEFT JOIN books ON books.id = borroweds.book_id")
	err := startQuery.Preload("User").Preload("Book").Limit(query.PerPage).Offset(offset).Find(&borroweds).Error
	if err != nil {
		return []models.Borrowed{}, 0, err
	}
	err = startQuery.Count(&count).Error
	if err != nil {
		return []models.Borrowed{}, 0, err
	}
	return borroweds, count, nil
}

func (r *BorrowedRepository) Create(borrowed models.Borrowed) (models.Borrowed, error) {
	err := r.db.Create(&borrowed).Error
	if err != nil {
		return models.Borrowed{}, err
	}
	return borrowed, nil
}

func (r *BorrowedRepository) Update(borrowed models.Borrowed) (models.Borrowed, error) {
	err := r.db.Save(&borrowed).Error
	if err != nil {
		return models.Borrowed{}, err
	}
	return borrowed, nil
}

func (r *BorrowedRepository) Delete(id int) error {
	return r.db.Delete(&models.Borrowed{}, id).Error
}
