package sqlite

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	DB *sqlx.DB
	// Auth *Auth
}

type SQLite interface {
	Close() error

	Account
	Auth
	Notes
	Project
	Section
	Task
	// Login(ctx context.Context, username string, password string) (*models.Account, error)
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

	db.SetMaxOpenConns(1)

	sqlite := &sqlite{
		DB: db,
	}

	err = sqlite.createAccountsTable()
	if err != nil {
		return nil, err
	}

	err = sqlite.createProjectsTable()
	if err != nil {
		return nil, err
	}

	err = sqlite.createProjectSectionsTable()
	if err != nil {
		return nil, err
	}

	err = sqlite.createTasksTable()
	if err != nil {
		return nil, err
	}

	err = sqlite.createDocumentsTable()
	if err != nil {
		return nil, err
	}

	err = sqlite.createDocumentSectionsTable()
	if err != nil {
		return nil, err
	}

	err = sqlite.createNotesTable()
	if err != nil {
		return nil, err
	}

	// err = sqlite.DBClose()

	return sqlite, nil
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

func (sqlite *sqlite) createAccountsTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		id TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in accounts table: %v", err)
	}
	return nil
}

func (sqlite *sqlite) createProjectsTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS projects (
		id TEXT NOT NULL,
		title TEXT NOT NULL,
		owner TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in projects table %v", err)
	}
	return nil
}

func (sqlite *sqlite) createDocumentsTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS documents (
		id TEXT NOT NULL,
		title TEXT NOT NULL,
		owner TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in documents table: %v", err)
	}
	return nil
}

func (sqlite *sqlite) createProjectSectionsTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS project_sections (
		id TEXT NOT NULL,
		project_id TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in project_sections table: %v", err)
	}
	return nil
}

func (sqlite *sqlite) createDocumentSectionsTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS document_sections (
		id TEXT NOT NULL,
		document_id TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in document_sections table: %v", err)
	}
	return nil
}

func (sqlite *sqlite) createTasksTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT NOT NULL,
		section_id TEXT NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in tasks table: %v", err)
	}
	return nil
}

func (sqlite *sqlite) createNotesTable() error {
	_, err := sqlite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS notes (
		id TEXT NOT NULL,
		section_id TEXT NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at TEXT NOT NULL,
		updated_at TEXT NOT NULL,
		PRIMARY KEY (id)
	)`)
	if err != nil {
		return fmt.Errorf("error in notes table")
	}
	return nil
}

func (sqlite *sqlite) Close() error {
	return sqlite.DB.Close()
}
