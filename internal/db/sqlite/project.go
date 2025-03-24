package sqlite

import "context"

type Project interface {
	// Project methods
	CreateProject(ctx context.Context) error
	DeleteProject(ctx context.Context) error
	GetProject(ctx context.Context) error
	UpdateProject(ctx context.Context) error
}

func (sqlite *sqlite) CreateProject(ctx context.Context) error {
	panic("Not completed")
}

func (sqlite *sqlite) DeleteProject(ctx context.Context) error {
	panic("Not completed")

}

func (sqlite *sqlite) GetProject(ctx context.Context) error {
	panic("Not completed")

}

func (sqlite *sqlite) UpdateProject(ctx context.Context) error {
	panic("Not completed")

}
