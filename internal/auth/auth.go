package auth

import (
	"context"
	"fmt"
	"time"

	"starlight/internal/db"
	"starlight/internal/db/sqlite"
	"starlight/internal/models"

	"github.com/google/uuid"
)

type auth struct {
	ctx    context.Context
	SQLite sqlite.SQLite
}

// type AuthConfig struct {
// 	sqliteRepo sqlite.SQLite
// }

func NewAuth(ctx context.Context, db *db.DB) *auth {
	return &auth{
		ctx:    ctx,
		SQLite: db.SQLite,
	}
}

func (auth *auth) CreateAccount(username string, password string) (bool, error) {
	if username == "" || password == "" {
		return false, fmt.Errorf("error: username and password must not be false")
	}
	newAccount := models.Account{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.UTC.String(),
		UpdatedAt: time.UTC.String(),
	}
	// fmt.Sprintln(newAccount)
	ctx := auth.ctx
	result, err := auth.SQLite.CreateAccount(ctx, &newAccount)
	if err != nil {
		panic("TODO")
	}
	return result, nil
}

func (auth *auth) Login(username string, password string) bool {

	hcUsername := "Crag Tarr"
	hcPassword := "Wails"

	if username == "" || password == "" || username != hcUsername || password != hcPassword {
		return false
	}

	if username == hcUsername && password == hcPassword {
		return true
	}

	return false
}

func (auth *auth) TestyLogin(username string, password string) string {
	result, err := auth.SQLite.TestyCreateAccount(username, password)
	if err != nil {
		return err.Error()
	}

	return result
}
