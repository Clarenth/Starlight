package repo

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"starlight/internal/models"
	// "github.com/mattn/go-sqlite3"
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

func SaveNoteToDB() {}

func SavePageToDB() {}

func SaveProjectToDB() {}

func SaveTaskToDB() {}

// func (c *DB) SaveUserToDB(username string, password string) {}

func (c *DB) SaveUserToDB(username string, password string) {
	newProfile := models.Account{
		ID:           strconv.Itoa(rand.Int()),
		Username:     username,
		Password:     password,
		ProfileImage: "temp",
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
	}

	c.memDB[newProfile.ID] = newProfile
}
