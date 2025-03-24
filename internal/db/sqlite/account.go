package sqlite

import (
	"context"
	"fmt"
	"log"

	"starlight/internal/models"
)

// Authentication
// func (sqlite *sqlite) Login(ctx context.Context, username string, password string) (*models.Account, error) {
// 	return nil, nil
// }

type Account interface {
	// Account methods
	CreateAccount(ctx context.Context, account *models.Account) (bool, error)
	DeleteAccount(ctx context.Context, accountID string) error
	GetAllAccounts(ctx context.Context) (*[]models.Account, error)
	UpdateAccount(ctx context.Context, username string, password string) error
}

// Account methods
func (sqlite *sqlite) CreateAccount(ctx context.Context, account *models.Account) (bool, error) {
	// return false, nil

	fmt.Printf("ID: '%v', Username '%v', Password: '%v', CreatedAt: '%v', UpdatedAt: '%v'", account.ID, account.Username, account.Password, account.CreatedAt, account.UpdatedAt)

	query := `INSERT OR IGNORE INTO accounts (id, username, password, created_at, updated_at) 
						VALUES ($1, $2, $3, $4, $5)
						RETURNING id, username, created_at, updated_at;
						`

	if err := sqlite.DB.Get(account, query, account.ID, account.Username, account.Password, account.CreatedAt, account.UpdatedAt); err != nil {
		log.Print(err)
		return false, fmt.Errorf("error: could not create account")
	}
	return true, nil
}

func (sqlite *sqlite) GetAllAccounts(ctx context.Context) (*[]models.Account, error) {
	accounts := &[]models.Account{}
	query := `SELECT id, username from accounts;`
	if err := sqlite.DB.SelectContext(ctx, accounts, query); err != nil {
		log.Printf("error in retriving AllAccounts from DB: %v", err)
		return nil, err
	}
	return accounts, nil
}

func (sqlite *sqlite) DeleteAccount(ctx context.Context, accountID string) error {
	deleteQuery := `DELETE FROM accounts WHERE id = $1`

	err := sqlite.DB.QueryRowContext(ctx, deleteQuery, accountID).Scan(&accountID)
	if err != nil {
		log.Printf("Error when deleting account with id %v", &accountID)
		return err
	}

	return nil
}

func (sqlite *sqlite) UpdateAccount(ctx context.Context, username string, password string) error {
	panic("Not completed")
}
