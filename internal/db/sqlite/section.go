package sqlite

import "context"

type Section interface {
	// Sections methods
	CreateSection(ctx context.Context) error
	GetSection(ctx context.Context) error
	DeleteSection(ctx context.Context) error
	UpdateSection(ctx context.Context) error
}

func (sqlite *sqlite) CreateSection(ctx context.Context) error {
	panic("Not completed")
}

func (sqlite *sqlite) GetSection(ctx context.Context) error {
	panic("Not completed")

}

func (sqlite *sqlite) DeleteSection(ctx context.Context) error {
	panic("Not completed")

}

func (sqlite *sqlite) UpdateSection(ctx context.Context) error {
	panic("Not completed")

}
