package models

import (
	"time"
	"victorubere/library/lib/types"
)

type Visitation struct {
	Model
	UserId    int            `json:"user_id"`
	VisitedAt time.Time      `json:"visited_at"`
	Duration  types.Duration `json:"duration"`
	User      User           `json:"user"`
}
