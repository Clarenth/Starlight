package sqlite

import (
	"log"
	"starlight/internal/models"
)

type Auth interface {
	GetPasswordHash(id string) (string, error)
}

func (sqlite *sqlite) GetPasswordHash(id string) (string, error) {
	hash := &models.Account{}
	query := `SELECT password FROM accounts WHERE id = $1;`

	if err := sqlite.DB.Get(hash, query, id); err != nil {
		log.Printf("error in SQLite GetPassword: %v", err)
		return err.Error(), err
	}

	return hash.Password, nil
}
