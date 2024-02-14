package sqlite

import "context"

func (sqlite *sqlite) CreateTask(ctx context.Context) error
func (sqlite *sqlite) GetTask(ctx context.Context) error
func (sqlite *sqlite) DeleteTask(ctx context.Context) error
func (sqlite *sqlite) UpdateTask(ctx context.Context) error
