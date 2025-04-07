package models

import (
	"victorubere/library/lib/structs"
	"victorubere/library/lib/types"
)

type BookRead struct {
	structs.Model
	UserID       uint           `json:"user_id"`
	BookID       uint           `json:"book_id"`
	VisitationID uint           `json:"visitation_id"`
	Duration     types.Duration `json:"duration"`
}
