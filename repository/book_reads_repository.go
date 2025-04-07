package repository

import (
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type BookReadsRepository struct {
	db *gorm.DB
}

func NewBookReadsRepository(db *gorm.DB) IBookReadssRepository {
	return &BookReadsRepository{
		db: db,
	}
}

func (r *BookReadsRepository) GetById(id int) (models.BookReads, error) {
	var bookRead models.BookReads
	err := r.db.Where("id = ?", id).Preload("User").Preload("Visitation").Preload("Book").First(&bookRead).Error
	if err != nil {
		return models.BookReads{}, err
	}
	return bookRead, nil
}

func (r *BookReadsRepository) List(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]models.BookReads, int64, error) {
	var bookReads []models.BookReads
	var count int64
	startQuery := r.db
	offset := helpers.GetOffset(query)
	if bookReadQuery.UserID != 0 {
		startQuery = startQuery.Where("user_id = ?", bookReadQuery.UserID)
	}
	if bookReadQuery.BookID != 0 {
		startQuery = startQuery.Where("book_id = ?", bookReadQuery.BookID)
	}
	if bookReadQuery.VisitationID != 0 {
		startQuery = startQuery.Where("visitation_id = ?", bookReadQuery.VisitationID)
	}
	if bookReadQuery.DurationStart != 0 {
		startQuery = startQuery.Where("duration >= ?", bookReadQuery.DurationStart)
	}
	if bookReadQuery.DurationEnd != 0 {
		startQuery = startQuery.Where("duration <= ?", bookReadQuery.DurationEnd)
	}
	startQuery = startQuery.Select("book_reads.*, book_reads.id as id").Joins("LEFT JOIN visitations ON visitations.id = book_reads.visitation_id").Joins("LEFT JOIN books ON books.id = book_reads.book_id").Joins("LEFT JOIN users ON users.id = book_reads.user_id")
	err := startQuery.Limit(query.PerPage).Offset(offset).Preload("User").Preload("Visitation").Preload("Visitation.User").Preload("Book").Find(&bookReads).Count(&count).Error
	if err != nil {
		return []models.BookReads{}, 0, err
	}
	
	return bookReads, count, nil
}

func (r *BookReadsRepository) Create(bookRead models.BookReads) (models.BookReads, error) {
	err := r.db.Create(&bookRead).Error
	if err != nil {
		return models.BookReads{}, err
	}
	return bookRead, nil
}

func (r *BookReadsRepository) Update(bookRead models.BookReads) (models.BookReads, error) {
	err := r.db.Save(&bookRead).Error
	if err != nil {
		return models.BookReads{}, err
	}
	return bookRead, nil
}

func (r *BookReadsRepository) Delete(id int) error {
	return r.db.Delete(&models.BookReads{}, id).Error
}
