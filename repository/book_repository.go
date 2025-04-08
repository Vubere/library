package repository

import (
	"time"
	"victorubere/library/lib/helpers"
	"victorubere/library/lib/structs"
	"victorubere/library/models"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) GetById(id int) (models.Book, error) {
	var book models.Book
	err := b.db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (b *BookRepository) List(query structs.Query, bookQuery structs.BookQuery) ([]models.Book, int64, error) {
	var books []models.Book
	var count int64
	dbExec := b.db
	offset := helpers.GetOffset(query)
	if bookQuery.Title != "" {
		dbExec = dbExec.Where("title LIKE ?", "%"+bookQuery.Title+"%")
	}
	if bookQuery.Author != "" {
		dbExec = dbExec.Where("author LIKE ?", "%"+bookQuery.Author+"%")
	}
	if bookQuery.Genre != "" {
		dbExec = dbExec.Where("genre = ?", bookQuery.Genre)
	}
	if bookQuery.ISBN != "" {
		dbExec = dbExec.Where("isbn = ?", bookQuery.ISBN)
	}
	if bookQuery.Publisher != "" {
		dbExec = dbExec.Where("publisher = ?", bookQuery.Publisher)
	}
	if bookQuery.Year != 0 {
		firstDayOfYear := time.Date(bookQuery.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
		lastDayOfYear := time.Date(bookQuery.Year, time.December, 31, 23, 59, 59, 0, time.UTC)
		dbExec = dbExec.Where("publication_date BETWEEN ? AND ?", firstDayOfYear, lastDayOfYear)
	}
	if bookQuery.BookYearsOld != 0 {
		timeYearsAgo := time.Now().AddDate(-bookQuery.BookYearsOld, 0, 0)
		dbExec = dbExec.Where("publication_date >= ?", timeYearsAgo)
	}
	err := dbExec.Limit(query.PerPage).Offset(offset).Find(&books).Error
	if err != nil {
		return []models.Book{}, 0, err
	}
	err = dbExec.Model(&books).Count(&count).Error
	if err != nil {
		return []models.Book{}, 0, err
	}
	return books, count, nil
}

func (b *BookRepository) Create(book models.Book) (models.Book, error) {
	err := b.db.Create(&book).Error
	if err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (b *BookRepository) Update(book models.Book) (models.Book, error) {
	var Book models.Book = book
	Book.UpdatedAt = time.Now()
	err := b.db.Model(&models.Book{}).Where("id = ?", Book.ID).Omit("id", "created_at", "updated_at").Updates(&Book).Error
	if err != nil {
		return models.Book{}, err
	}
	return Book, nil
}

func (b *BookRepository) Delete(id int) error {
	return b.db.Delete(&models.Book{}, id).Error
}

func (b *BookRepository) TotalBooks(bookQuery structs.BookQuery) (int64, error) {
	var count int64
	startQuery := b.db.Model(&models.Book{})
	if bookQuery.Title != "" {
		startQuery = startQuery.Where("title LIKE ?", "%"+bookQuery.Title+"%")
	}
	if bookQuery.Author != "" {
		startQuery = startQuery.Where("author LIKE ?", "%"+bookQuery.Author+"%")
	}
	if bookQuery.Genre != "" {
		startQuery = startQuery.Where("genre = ?", bookQuery.Genre)
	}
	if bookQuery.ISBN != "" {
		startQuery = startQuery.Where("isbn = ?", bookQuery.ISBN)
	}
	if bookQuery.Publisher != "" {
		startQuery = startQuery.Where("publisher = ?", bookQuery.Publisher)
	}
	if bookQuery.Year != 0 {
		firstDayOfYear := time.Date(bookQuery.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
		lastDayOfYear := time.Date(bookQuery.Year, time.December, 31, 23, 59, 59, 0, time.UTC)
		startQuery = startQuery.Where("publication_date BETWEEN ? AND ?", firstDayOfYear, lastDayOfYear)
	}
	if bookQuery.BookYearsOld != 0 {
		timeYearsAgo := time.Now().AddDate(-bookQuery.BookYearsOld, 0, 0)
		startQuery = startQuery.Where("publication_date >= ?", timeYearsAgo)
	}
	err := startQuery.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
