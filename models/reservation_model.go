package models

import (
	"time"
	"victorubere/library/lib/types"

	"gorm.io/gorm"
)

type Reservation struct {
	UserId        int
	BookId        int
	ReservationAt time.Time
	Duration      types.Duration
}

type ReservationModel struct {
	gorm.Model
	Reservation
}

type ReservationJSON struct {
	UserId        int            `json:"user_id"`
	BookId        int            `json:"book_id"`
	ReservationAt string         `json:"reservation_at"`
	ReturnedAt    string         `json:"returned_at"`
	Duration      types.Duration `json:"duration"`
}
