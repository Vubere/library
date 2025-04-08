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

func NewBookReadsRepository(db *gorm.DB) IBookReadsRepository {
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
	err := startQuery.Limit(query.PerPage).Offset(offset).Preload("User").Preload("Visitation").Preload("Visitation.User").Preload("Book").Find(&bookReads).Error
	if err != nil {
		return []models.BookReads{}, 0, err
	}
	err = startQuery.Model(&bookReads).Count(&count).Error
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

func (r *BookReadsRepository) TotalBookReads(bookReadQuery structs.BookReadsQuery) (int64, error) {
	var count int64
	startQuery := r.db.Model(&models.BookReads{})
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
	err := startQuery.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *BookReadsRepository) MostReadBooks(query structs.Query, bookReadQuery structs.BookReadsQuery) ([]structs.MostBookReadsDTO, error) {
	var bookReads []structs.MostBookReadsDTO
	startQuery := r.db.Model(&models.BookReads{})
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
	startQuery = startQuery.Joins("LEFT JOIN books on books.id = book_reads.book_id").Select("count(books.id) as book_reads_count, books.id as book_id, books.title as book_title, books.author as book_author, books.isbn as book_isbn, books.publisher as book_publisher, books.publication_date as book_publication_date").Group("book_id").Order("book_reads_count desc")
	if query.PerPage != 0 && query.Page != 0 {
		startQuery = startQuery.Offset((query.Page - 1) * query.PerPage).Limit(query.PerPage)
	}
	err := startQuery.Find(&bookReads).Error
	if err != nil {
		return []structs.MostBookReadsDTO{}, err
	}
	return bookReads, nil
}

func (r *BookReadsRepository) UserWithMostBookReads(bookId int, bookReadQuery structs.BookReadsQuery) (structs.BookReadMostByUserDTO, error) {
	var bookReads structs.BookReadMostByUserDTO
	startQuery := r.db.Model(&models.BookReads{}).Where("book_id = ?", bookId)

	if bookReadQuery.VisitationID != 0 {
		startQuery = startQuery.Where("visitation_id = ?", bookReadQuery.VisitationID)
	}
	if bookReadQuery.DurationStart != 0 {
		startQuery = startQuery.Where("duration >= ?", bookReadQuery.DurationStart)
	}
	if bookReadQuery.DurationEnd != 0 {
		startQuery = startQuery.Where("duration <= ?", bookReadQuery.DurationEnd)
	}
	err := startQuery.Joins("LEFT JOIN users on users.id = book_reads.user_id").Select("count(book_reads.user_id) as book_reads_count, book_reads.user_id as user_id, users.name as user_name, users.gender as user_gender, users.date_of_birth as user_date_of_birth, users.email as user_email").Group("user_id").Order("book_reads_count desc").Find(&bookReads).Error
	if err != nil {
		return structs.BookReadMostByUserDTO{}, err
	}
	return bookReads, nil
}
