package sqlite

import (
	"context"
	"fmt"
	"starlight/internal/models"
	"time"

	"github.com/google/uuid"
)

type Account interface {
	CreateAccount(ctx context.Context, username string, password string)
}

func (sqlite sqlite) createAccountTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		id TEXT NOT NULL
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in accounts table!")
	}
	return nil
}

func (sqlite *sqlite) CreateAccount(ctx context.Context, username string, password string) (string, error) {
	account := models.Account{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	query := `INSERT OR IGNORE INTO accounts (id, username, password) 
						VALUES ($1, $2, $3)
						`

	if err := sqlite.DB.Get(account, query); err != nil {
		return "", fmt.Errorf("error: could not create account")
	}
	return fmt.Sprintf("created account with username %v", username), nil
}
