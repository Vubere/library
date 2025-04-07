package models

import (
	"time"
	"victorubere/library/lib/structs"
	"victorubere/library/lib/types"
)

type Visitations struct {
	structs.Model
	UserId    int            `json:"user_id"`
	VisitedAt time.Time      `json:"visited_at"`
	Duration  types.Duration `json:"duration"`
}
