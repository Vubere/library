package models

import (
	"time"
)

type Borrowed struct {
	Model
	UserId     int       `json:"user_id"`
	BookId     int       `json:"book_id"`
	BorrowedAt time.Time `json:"borrowed_at"`
	ReturnedAt time.Time `json:"returned_at"`
	User       User      `json:"user"`
	Book       Book      `json:"book"`
}
