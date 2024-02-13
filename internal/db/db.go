package db

import (
	"fmt"
	"log"

	"starlight/internal/db/sqlite"
)

type DB struct {
	SQLite sqlite.SQLite
}

func NewDB() *DB {
	sqlite, err := sqlite.NewSqlite()
	if err != nil {
		log.Print(fmt.Errorf("error initialising SQLite DB: %v", err))
	}

	return &DB{
		SQLite: sqlite,
	}
}

// func (db *DB) CreateAccount(username string, password string) (string, error) {
// 	result, err := db.SQLite.CreateAccount()
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }
