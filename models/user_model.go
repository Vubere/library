package models

import (
	"time"
	"victorubere/library/lib/structs"
)


type User struct {
	structs.Model
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address    string    `json:"address"`
	Gender     string    `json:"gender"`
	Role        string    `json:"role"`
}

