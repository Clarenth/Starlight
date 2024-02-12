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

func NewAccount() *Account {
	return &Account{}
}

type Note struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type Page struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type Project struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Owner     string    `json:"owner"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
