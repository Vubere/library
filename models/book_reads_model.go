package models

import (
	"victorubere/library/lib/types"
)

type BookReads struct {
	Model
	UserID       uint           `json:"user_id"`
	BookID       uint           `json:"book_id"`
	VisitationID uint           `json:"visitation_id"`
	Duration     types.Duration `json:"duration"`
	User         User           `json:"user"`
	Visitation   Visitation     `json:"visitation"`
	Book         Book           `json:"book"`
}
