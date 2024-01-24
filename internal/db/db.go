package db

import (
	"context"
	"fmt"

	"starlight/internal/models"
	// sqlite3 "github.com/mattn/go-sqlite3"
)

type DB struct {
	ctx   context.Context
	memDB map[string]models.Account
}

func NewDB() *DB {
	return &DB{
		memDB: make(map[string]models.Account),
	}
}

func (c *DB) SetContext(ctx context.Context) {
	c.ctx = ctx
}

func (c *DB) newSQLiteDB() {
}

/*Old Test Functions*/
func (c *DB) NewMemDB() map[string]models.Account {
	newDB := make(map[string]models.Account, 5)
	return newDB
}

func (c *DB) PrintDB() string {
	return fmt.Sprint(c.memDB)
}

func (db *DB) LoadAccountData(name string) string {
	output := fmt.Sprintf("Hello %v!", name)
	return output
}
