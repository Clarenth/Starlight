package sqlite

import (
	"context"
	"fmt"
	"starlight/internal/models"
	"time"

	"github.com/google/uuid"
)

func (sqlite *sqlite) createAccountTable() error {
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

	query := `INSERT OR IGNORE INTO accounts (id, username, password, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5)
						`

	if err := sqlite.DB.Get(account, query); err != nil {
		return "", fmt.Errorf("error: could not create account")
	}
	return fmt.Sprintf("created account with username %v", username), nil
}

func (sqlite *sqlite) TestyCreateAccount(username string, password string) (string, error) {
	account := models.Account{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	if (models.Account{} == account) {
		return "", fmt.Errorf("zero value was detected in TestyCreateAccount")
	}
	return fmt.Sprint("Hello from TestyCreateAccount"), nil
}
