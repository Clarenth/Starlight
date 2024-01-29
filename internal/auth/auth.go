package auth

import (
	"fmt"
	"starlight/internal/models"
	"time"

	"github.com/google/uuid"
)

type auth struct{}

func New() *auth {
	return &auth{}
}

func (*auth) CreateAccount(username string, password string) (string, bool) {
	if username == "" || password == "" {
		return "error: username and password must not be false", false
	}
	newAccount := &models.Account{
		ID:        uuid.New().String(),
		Username:  username,
		Password:  password,
		CreatedAt: time.UTC.String(),
		UpdatedAt: time.UTC.String(),
	}
	return fmt.Sprint(newAccount), true
}

func (*auth) Login(username string, password string) bool {

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
