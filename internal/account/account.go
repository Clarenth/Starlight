package account

import (
	"context"
	"log"

	"starlight/internal/db"
	"starlight/internal/db/sqlite"
	"starlight/internal/models"
)

type account struct {
	ctx    context.Context
	SQLite sqlite.SQLite
}

func NewAccount(ctx context.Context, db *db.DB) *account {
	return &account{
		ctx:    ctx,
		SQLite: db.SQLite,
	}
}

func (account *account) GetAccounts() (*[]models.Account, error) {
	accounts, err := account.SQLite.GetAllAccounts(account.ctx)
	if err != nil {
		log.Printf("error in GetAccounts: %v", err)
		return nil, err
	}
	return accounts, nil
}
