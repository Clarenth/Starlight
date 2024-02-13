package auth

import (
	"fmt"
	"time"

	"starlight/internal/db"
	"starlight/internal/db/sqlite"
	"starlight/internal/models"

	"github.com/google/uuid"
)

type auth struct {
	DB sqlite.SQLite
}

// type AuthConfig struct {
// 	sqliteRepo sqlite.SQLite
// }

func NewAuth(db *db.DB) *auth {
	return &auth{
		DB: db.SQLite,
	}
}

func (auth *auth) CreateAccount(username string, password string) (string, bool) {
	if username == "" || password == "" {
		return "error: username and password must not be false", false
	}
	newAccount := models.Account{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.UTC.String(),
		UpdatedAt: time.UTC.String(),
	}
	fmt.Printf("%v", newAccount)
	panic("not done yet!")
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
	result, err := auth.DB.TestyCreateAccount(username, password)
	if err != nil {
		return err.Error()
	}

	return result
}
