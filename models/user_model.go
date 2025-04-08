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
	Password   	string  	`json:"password,omitempty"`
}

type UserDTO struct {
	ID          uint      `json:"id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Address     string    `json:"address,omitempty"`
	Gender      string    `json:"gender,omitempty"`
	Role        string    `json:"role,omitempty"`
}
