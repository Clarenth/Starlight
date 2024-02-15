package db

import (
	"context"
	"fmt"
	"log"

	"starlight/internal/db/sqlite"
)

type DB struct {
	ctx    context.Context
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
