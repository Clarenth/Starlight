package account

import (
	"context"
	"errors"
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

func (account *account) DeleteAccount(id string) error {
	if uuid.Validate(id) != nil {
		return errors.New("id is invalid")
	}

	err := account.SQLite.DeleteAccount(id)
	if err != nil {
		return errors.New("unable to delete account")
	}

	return nil
}

func (account *account) GetOneAccount(id uuid.UUID) *models.Account {
	idString := id.String()
	data, err := account.SQLite.GetAccountData(idString)
	if err != nil {
		log.Panicf("blah")
		return nil
	}
	return data
}

func (account *account) GetAllAccounts() *[]models.Account {
	accounts, err := account.SQLite.GetAllAccounts()
	if err != nil {
		log.Printf("error in GetAccounts: %v", err)
		return nil
	}
	return accounts
}
