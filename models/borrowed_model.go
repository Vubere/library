package models

import (
	"victorubere/library/lib/structs"
)

type Borrowed struct {
	structs.Model
	UserId        int            `json:"user_id"`
	BookId        int            `json:"book_id"`
	BorrowedAt string         `json:"borrowed_at"`
	ReturnedAt    string         `json:"returned_at"`
}
