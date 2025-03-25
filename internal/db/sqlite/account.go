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
	DeleteAccount(accountID string) error
	GetAccountData(id string) (*models.Account, error)
	GetAllAccounts() (*[]models.Account, error)
	UpdateAccount(ctx context.Context, username string, password string) error
}

// Account methods
func (sqlite *sqlite) CreateAccount(ctx context.Context, account *models.Account) (bool, error) {
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

func (sqlite *sqlite) DeleteAccount(accountID string) error {
	deleteQuery := `DELETE FROM accounts WHERE id = $1`

	// Add error handling for if an invalid ID is passed such as
	// a non-existent ID in the Database

	err := sqlite.DB.QueryRow(deleteQuery, accountID).Scan(&accountID)
	if err != nil {
		log.Printf("Error when deleting account with id %v", &accountID)
		return err
	}

	return nil
}

// Starlight: look into why CTX in account and likely auth are nil. Why is there no ctx?
func (sqlite *sqlite) GetAllAccounts() (*[]models.Account, error) {
	accounts := &[]models.Account{}
	query := `SELECT id, username, created_at, updated_at FROM accounts;`
	// query := `SELECT
	// 								id,
	// 								username,
	// 								to_char(created_at, 'yyyy-mm-dd') AS created_at,
	// 								to_char(updated_at, 'yyyy-mm-dd') AS updated_at,
	// 					FROM accounts;`
	if err := sqlite.DB.Select(accounts, query); err != nil {
		log.Printf("error in retriving AllAccounts from DB: %v", err)
		return nil, err
	}
	return accounts, nil
}

// Starlight: look into why CTX in account and likely auth are nil. Why is there no ctx?
func (sqlite *sqlite) GetAccountData(id string) (*models.Account, error) {
	account := &models.Account{}
	query := `SELECT id, username, created_at, updated_at FROM accounts WHERE id = $1;`

	if err := sqlite.DB.Get(account, query, id); err != nil {
		log.Printf("error in GetAccountData: %v", err)
		return nil, err
	}

	return account, nil
}

func (sqlite *sqlite) UpdateAccount(ctx context.Context, username string, password string) error {
	panic("Not completed")
}
