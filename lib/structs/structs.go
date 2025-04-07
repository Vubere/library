package structs

import (
	"time"
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
	UserID    int `form:"user_id"`
	VisitedAt time.Time `form:"visited_at"`
	Duration  int `form:"duration"`
	VisitedAtStart time.Time `form:"visited_at_start"`
	VisitedAtEnd time.Time `form:"visited_at_end"`
}

type BookReadQuery struct {
	UserID    int `form:"user_id"`
	BookID    int `form:"book_id"`
	VisitationID int `form:"visitation_id"`
	DurationStart int `form:"duration_start"`
	DurationEnd int `form:"duration_end"`
}

type BorrowedQuery struct {
	UserID    int `form:"user_id"`
	BookID    int `form:"book_id"`
	Duration  int `form:"duration"`
	BorrowedAtStart time.Time `form:"borrowed_at_start"`
	BorrowedAtEnd time.Time `form:"borrowed_at_end"`
	ReturnedAtStart time.Time `form:"returned_at_start"`
	ReturnedAtEnd time.Time `form:"returned_at_end"`
}

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Meta struct {
	TotalRecords      int64 `json:"total_records"`
	Page              int   `json:"page"`
	PerPage           int   `json:"per_page"`
	NextPage          int   `json:"next_page"`
	NextPageAvailable bool  `json:"next_page_available"`
}
