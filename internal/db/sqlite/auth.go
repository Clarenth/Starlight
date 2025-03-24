package sqlite

import (
	"context"
	"starlight/internal/models"
)

type Auth interface {
	Login(ctx context.Context, username string, password string) (*models.Account, error)
}

// Authentication
func (sqlite *sqlite) Login(ctx context.Context, username string, password string) (*models.Account, error) {
	return nil, nil
}
