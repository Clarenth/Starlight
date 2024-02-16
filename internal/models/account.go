package models

import "github.com/google/uuid"

type Account struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	ProfileImage string    `json:"profile_image"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
}
