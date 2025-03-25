package sqlite

import (
	"context"
	"log"
	"starlight/internal/models"
)

type Auth interface {
	GetPasswordHash(ctx context.Context, id string) (string, error)
}

func (sqlite *sqlite) GetPasswordHash(ctx context.Context, id string) (string, error) {
	hash := &models.Account{}
	query := `SELECT password FROM accounts WHERE id = $1;`

	if err := sqlite.DB.GetContext(ctx, hash, query, id); err != nil {
		log.Printf("error in SQLite GetPassword: %v", err)
		return err.Error(), err
	}

	return hash.Password, nil
}
