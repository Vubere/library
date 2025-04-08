package models

import (
	"time"
)

type User struct {
	Model
	Name        string    `json:"name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender"`
	Role        string    `json:"role"`
	Password   	string  	`json:"password"`
}

type UserDTO struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Gender      string    `json:"gender"`
	Role        string    `json:"role"`
}
