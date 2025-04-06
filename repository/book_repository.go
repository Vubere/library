package repository

import (
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
	if bookQuery.Year != "" {
		dbExec = dbExec.Where("year = ?", bookQuery.Year)
	}
	err := dbExec.Limit(query.PerPage).Offset(offset).Find(&books).Count(&count).Error
	if err != nil {
		return []models.Book{}, 0, err
	}
	return books, count, nil
}

func (b *BookRepository) Create(book models.Book) (models.Book, error) {
	return models.Book{}, nil
}

func (b *BookRepository) Update(book models.Book) (models.Book, error) {
	return models.Book{}, nil
}

func (b *BookRepository) Delete(id int) error {
	return b.db.Delete(&models.Book{}, id).Error
}
