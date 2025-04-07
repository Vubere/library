package repository

import (
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type BookReadRepository struct {
	db *gorm.DB
}

func NewBookReadRepository(db *gorm.DB) IBookReadsRepository {
	return &BookReadRepository{
		db: db,
	}
}

func (r *BookReadRepository) GetById(id int) (models.BookRead, error) {
	var bookRead models.BookRead
	err := r.db.Where("id = ?", id).First(&bookRead).Error
	if err != nil {
		return models.BookRead{}, err
	}
	return bookRead, nil
}

func (r *BookReadRepository) List(query structs.Query, bookReadQuery structs.BookReadQuery) ([]models.BookRead, int64, error) {
	var bookReads []models.BookRead
	var count int64
	dbExec := r.db
	offset := helpers.GetOffset(query)
	if bookReadQuery.UserID != 0 {
		dbExec = dbExec.Where("user_id = ?", bookReadQuery.UserID)
	}
	if bookReadQuery.BookID != 0 {
		dbExec = dbExec.Where("book_id = ?", bookReadQuery.BookID)
	}
	if bookReadQuery.VisitationID != 0 {
		dbExec = dbExec.Where("visitation_id = ?", bookReadQuery.VisitationID)
	}
	if bookReadQuery.DurationStart != 0 {
		dbExec = dbExec.Where("duration >= ?", bookReadQuery.DurationStart)
	}
	if bookReadQuery.DurationEnd != 0 {
		dbExec = dbExec.Where("duration <= ?", bookReadQuery.DurationEnd)
	}
	err := dbExec.Limit(query.PerPage).Offset(offset).Find(&bookReads).Count(&count).Error
	if err != nil {
		return []models.BookRead{}, 0, err
	}
	return bookReads, count, nil
}

func (r *BookReadRepository) Create(bookRead models.BookRead) (models.BookRead, error) {
	err := r.db.Create(&bookRead).Error
	if err != nil {
		return models.BookRead{}, err
	}
	return bookRead, nil
}

func (r *BookReadRepository) Update(bookRead models.BookRead) (models.BookRead, error) {
	err := r.db.Save(&bookRead).Error
	if err != nil {
		return models.BookRead{}, err
	}
	return bookRead, nil
}

func (r *BookReadRepository) Delete(id int) error {
	return r.db.Delete(&models.BookRead{}, id).Error
}