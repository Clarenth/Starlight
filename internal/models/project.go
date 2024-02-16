package models

import "github.com/google/uuid"

type Project struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Owner     string    `json:"owner"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type ProjecSection struct {
	ID        uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type Task struct {
	ID          uuid.UUID `json:"id"`
	SectionID   uuid.UUID `json:"section_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
