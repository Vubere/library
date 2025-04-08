package structs

import (
	"time"
	"victorubere/library/models"
)

type Pagination struct {
	Skip  int
	Limit int
}

func NewPagination(page, PerPage int) Pagination {
	p := Pagination{}
	p.Skip = (page - 1) * PerPage
	p.Limit = PerPage
	return p
}

type Sort struct {
	Field string
	Order string // ASC or DESC
}

type Query struct {
	Page          int    `form:"page"`
	PerPage       int    `form:"per_page"`
	SortBy        string `form:"sort_by"`
	SortDirection string `form:"sort_direction"`
}

type UserQuery struct {
	Email            string `form:"email"`
	Name             string `form:"name"`
	Gender           string `form:"gender"`
	DateCreatedEnd   string `form:"date_created_end"`
	DateCreatedStart string `form:"date_created_start"`
	MinAge           int    `form:"min_age"`
}

type BookQuery struct {
	Title        string `form:"title"`
	Author       string `form:"author"`
	Genre        string `form:"genre"`
	ISBN         string `form:"isbn"`
	Publisher    string `form:"publisher"`
	BookYearsOld int    `form:"book_years_old"`
	Year         int    `form:"year"`
}

type VisitationQuery struct {
	UserID         int       `form:"user_id"`
	VisitedAt      time.Time `form:"visited_at"`
	Duration       int       `form:"duration"`
	VisitedAtStart time.Time `form:"visited_at_start"`
	VisitedAtEnd   time.Time `form:"visited_at_end"`
}

type BookReadsQuery struct {
	UserID        int `form:"user_id"`
	BookID        int `form:"book_id"`
	VisitationID  int `form:"visitation_id"`
	DurationStart int `form:"duration_start"`
	DurationEnd   int `form:"duration_end"`
}

type BorrowedQuery struct {
	UserID          int       `form:"user_id"`
	BookID          int       `form:"book_id"`
	Duration        int       `form:"duration"`
	BorrowedAtStart time.Time `form:"borrowed_at_start"`
	BorrowedAtEnd   time.Time `form:"borrowed_at_end"`
	ReturnedAtStart time.Time `form:"returned_at_start"`
	ReturnedAtEnd   time.Time `form:"returned_at_end"`
}

type Meta struct {
	TotalRecords      int64 `json:"total_records"`
	Page              int   `json:"page"`
	PerPage           int   `json:"per_page"`
	NextPage          int   `json:"next_page"`
	NextPageAvailable bool  `json:"next_page_available"`
}

type UserSummaryDTO struct {
	UserDetails         models.UserDTO         `json:"user_details"`
	VisitationsCount    int64               `json:"visitations_count"`
	BorrowedsCount      int64               `json:"borroweds_count"`
	BookReadsCount      int64               `json:"book_reads_count"`
	MostReadBook        *MostBookReadsDTO    `json:"most_read_book"`
	MostBorrowedBook	  *MostBorrowedBookDTO `json:"most_borrowed_book"`
}

type BooksSummaryDTO struct {
	BooksCount          int64               `json:"books_count"`
	MostReadBook        *MostBookReadsDTO    `json:"most_read_book"`
	MostBorrowedBook *MostBorrowedBookDTO `json:"most_borrowed_book"`
}

type BookSummaryDTO struct {
	BookDetails           models.Book           `json:"book_details"`
	ReadsCount            int64                 `json:"reads_count"`
	BorrowedsCount        int64                 `json:"borroweds_count"`
	ReadMostByUser        *BookReadMostByUserDTO `json:"read_most_by_user"`
	BorrowedMostByUser *BorrowedMostByUserDTO `json:"borrowed_most_by_user"`
}

type MostBookReadsDTO struct {
	BookReadsCount      int    `json:"book_reads_count"`
	BookID              int    `json:"book_id"`
	BookTitle           string `json:"book_title"`
	BookAuthor          string `json:"book_author"`
	BookGenre           string `json:"book_genre"`
	BookISBN            string `json:"book_isbn"`
	BookPublisher       string `json:"book_publisher"`
	BookPublicationDate string `json:"book_publication_date"`
	BookYear            int    `json:"book_year"`
}

type MostBorrowedBookDTO struct {
	BookBorrowedCount   int    `json:"book_borrowing_count"`
	BookID              int    `json:"book_id"`
	BookTitle           string `json:"book_title"`
	BookAuthor          string `json:"book_author"`
	BookGenre           string `json:"book_genre"`
	BookISBN            string `json:"book_isbn"`
	BookPublisher       string `json:"book_publisher"`
	BookPublicationDate string `json:"book_publication_date"`
	BookYear            int    `json:"book_year"`
}

type BookReadMostByUserDTO struct {
	BookReadsCount  int    `json:"book_reads_count"`
	UserID          int    `json:"user_id"`
	UserName        string `json:"user_name"`
	UserGender      string `json:"user_gender"`
	UserDateOfBirth string `json:"user_date_of_birth"`
	UserEmail       string `json:"user_email"`
}

type BorrowedMostByUserDTO struct {
	UserBorrowingCount int    `json:"user_borrowing_count"`
	UserID             int    `json:"user_id"`
	UserName           string `json:"user_name"`
	UserGender         string `json:"user_gender"`
	UserDateOfBirth    string `json:"user_date_of_birth"`
	UserEmail          string `json:"user_email"`
}
