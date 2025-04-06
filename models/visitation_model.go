package models

import (
	"time"
	"victorubere/library/lib/types"

	"gorm.io/gorm"
)

type Visitation struct {
	UserId    int
	VisitedAt time.Time
	Duration  types.Duration
}

type VisitationModel struct {
	gorm.Model
	Visitation
}

type VisitationJSON struct {
	UserId    int            `json:"user_id"`
	VisitedAt time.Time      `json:"visited_at"`
	Duration  types.Duration `json:"duration"`
}
