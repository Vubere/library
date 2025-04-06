package models

import (
	"time"
	"victorubere/library/lib/structs"
)

type Borrowed struct {
	structs.Model
	UserId        int            `json:"user_id"`
	BookId        int            `json:"book_id"`
	BorrowedAt 		time.Time         `json:"borrowed_at"`
	ReturnedAt    time.Time         `json:"returned_at"`
}
