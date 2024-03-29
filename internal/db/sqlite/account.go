package sqlite

import (
	"context"
	"fmt"
	"time"

	"starlight/internal/models"

	"github.com/google/uuid"
)

func (sqlite *sqlite) CreateAccount(ctx context.Context, account *models.Account) (bool, error) {
	return false, nil

	query := `INSERT OR IGNORE INTO accounts (id, username, password, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5)
						`

	if err := sqlite.DB.Get(account, query); err != nil {
		return false, fmt.Errorf("error: could not create account")
	}
	return true, nil
}

func (sqlite *sqlite) GetAccount(ctx context.Context, username string, password string) (*models.Account, error) {
	// temp info
	account := &models.Account{
		ID:        uuid.New(),
		Username:  username,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	return account, nil
	panic("Not completed")
}

func (sqlite *sqlite) DeleteAccount(ctx context.Context, username string, password string) error {
	panic("Not completed")
}

func (sqlite *sqlite) UpdateAccount(ctx context.Context, username string, password string) error {
	panic("Not completed")
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
