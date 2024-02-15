package sqlite

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"starlight/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	DB *sqlx.DB
}

type SQLite interface {
	CreateAccount(ctx context.Context, account *models.Account) (string, error)
	DeleteAccount(ctx context.Context, username string, password string) error
	GetAccount(ctx context.Context, username string, password string) error
	UpdateAccount(ctx context.Context, username string, password string) error
	TestyCreateAccount(username string, password string) (string, error)

	CreateNote(ctx context.Context) error
	DeleteNote(ctx context.Context) error
	GetNote(ctx context.Context) error
	UpdateNote(ctx context.Context) error

	CreateProject(ctx context.Context) error
	DeleteProject(ctx context.Context) error
	GetProject(ctx context.Context) error
	UpdateProject(ctx context.Context) error

	CreateSection(ctx context.Context) error
	GetSection(ctx context.Context) error
	DeleteSection(ctx context.Context) error
	UpdateSection(ctx context.Context) error

	CreateTask(ctx context.Context) error
	DeleteTask(ctx context.Context) error
	GetTask(ctx context.Context) error
	UpdateTask(ctx context.Context) error
}

func NewSqlite() (SQLite, error) {
	sqliteFile, err := initDataSource()
	if err != nil {
		log.Fatal("error: could not initalise SQLite db file")
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

func (sqlite *sqlite) createAccountTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		id TEXT NOT NULL
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in accounts table")
	}
	return nil
}
