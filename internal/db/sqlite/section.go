package sqlite

import "context"

func (sqlite *sqlite) CreateSection(ctx context.Context) error
func (sqlite *sqlite) GetSection(ctx context.Context) error
func (sqlite *sqlite) DeleteSection(ctx context.Context) error
func (sqlite *sqlite) UpdateSection(ctx context.Context) error
