package sqlite

import "context"

type Task interface {
	CreateTask(ctx context.Context) error
	DeleteTask(ctx context.Context) error
	GetTask(ctx context.Context) error
	UpdateTask(ctx context.Context) error
}

func (sqlite *sqlite) CreateTask(ctx context.Context) error {
	panic("Not completed")
}

func (sqlite *sqlite) GetTask(ctx context.Context) error {
	panic("Not completed")

}

func (sqlite *sqlite) DeleteTask(ctx context.Context) error {
	panic("Not completed")

}

func (sqlite *sqlite) UpdateTask(ctx context.Context) error {
	panic("Not completed")

}
