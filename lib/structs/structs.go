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
