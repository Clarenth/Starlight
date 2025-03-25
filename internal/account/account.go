package account

import (
	"context"
	"log"

	"starlight/internal/db"
	"starlight/internal/db/sqlite"
	"starlight/internal/models"

	"github.com/google/uuid"
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

func (account *account) GetOneAccount(id uuid.UUID) (*models.Account, error) {
	idString := id.String()
	data, err := account.SQLite.GetAccountData(idString)
	if err != nil {
		log.Panicf("blah")
		return nil, err
	}
	return data, nil
}

func (account *account) GetAccounts() *[]models.Account {
	accounts, err := account.SQLite.GetAllAccounts()
	if err != nil {
		log.Printf("error in GetAccounts: %v", err)
		return nil
	}
	return accounts
}
