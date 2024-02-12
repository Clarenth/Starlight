package sqlite

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	DB *sqlx.DB
}

type SQLite interface {
	CreateAccount(ctx context.Context, username string, password string) (string, error)
}

func NewSqlite() (SQLite, error) {
	sqliteFile, err := initDataSource()
	if err != nil {
		return nil, fmt.Errorf("error: could not initalise SQLite db file")
	}

	db, err := sqlx.Open("sqlite3", sqliteFile)
	if err != nil {
		return nil, fmt.Errorf("sqlite failed to open: %w", err)
	}

	return &sqlite{
		DB: db,
	}, nil
}

func initDataSource() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "error in user home directory: Starlight failed to start SQLite", err
	}
	dataDir := filepath.Join(homeDir, "starlight")
	os.Mkdir(dataDir, os.FileMode(0755))
	return filepath.Join(dataDir, "starlight.db"), nil
}
