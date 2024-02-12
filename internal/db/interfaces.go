package db

import "context"

type SQLite interface {
	CreateAccount(ctx context.Context, username string, password string) (string, error)
}
