package sqlite

import "context"

type Notes interface {
	// Notes methods
	CreateNote(ctx context.Context) error
	DeleteNote(ctx context.Context) error
	GetNote(ctx context.Context) error
	UpdateNote(ctx context.Context) error
}

func (sqlite *sqlite) CreateNote(ctx context.Context) error {
	panic("Not completed yet")
}

func (sqlite *sqlite) DeleteNote(ctx context.Context) error {
	panic("Not completed yet")
}

func (sqlite *sqlite) GetNote(ctx context.Context) error {
	panic("Not completed yet")
}

func (sqlite *sqlite) UpdateNote(ctx context.Context) error {
	panic("Not completed yet")
}
