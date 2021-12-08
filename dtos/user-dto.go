package dtos

import "time"

type UserDTO struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
