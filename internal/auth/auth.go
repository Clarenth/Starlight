package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"starlight/internal/db"
	"starlight/internal/db/sqlite"
	"starlight/internal/models"

	"github.com/alexedwards/argon2id"
	"github.com/google/uuid"
)

type auth struct {
	ctx    context.Context
	SQLite sqlite.SQLite
}

var argon2Params = &argon2id.Params{
	Memory:      62500,
	Iterations:  2,
	Parallelism: 16,
	SaltLength:  32,
	KeyLength:   64,
}

// type AuthConfig struct {
// 	sqliteRepo sqlite.SQLite
// }

func NewAuth(ctx context.Context, db *db.DB) *auth {
	return &auth{
		ctx:    ctx,
		SQLite: db.SQLite,
	}
}

func (auth *auth) CreateAccount(username string, password string) (bool, error) {
	if username == "" || password == "" {
		return false, fmt.Errorf("error: username and password must not be false")
	}

	// implement argon2 hashing and salting here
	// to protect passwords from plain text exposing
	hashedPassword, err := argon2id.CreateHash(password, argon2Params)
	if err != nil {
		log.Panicf("Unable to encrypt new user password in CreateAccount: %v\n", username)
		return false, err
	}
	password = hashedPassword //The hashed password is saved to the DB instead of the plain text. User Signin will compare the client's input password to the stored hash and validity

	newAccount := models.Account{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	}
	fmt.Sprintln(newAccount)
	ctx := auth.ctx
	result, err := auth.SQLite.CreateAccount(ctx, &newAccount)
	if err != nil {
		panic("Failed to create account.")
	}
	return result, nil
}

func (auth *auth) Login(id string, username string, password string) (*models.Account, error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("username and password cannot be empty")
	}

	if uuid.Validate(id) != nil {
		return nil, errors.New("profile ID is invalid")
	}

	// idString := id.String()

	// Here this will take in the user ID and password. The ID and Username will be sent to the frontend
	// when the app is booted up for the first time (if it has not been opened or closed that day)
	hash, err := auth.SQLite.GetPasswordHash(id)
	if err != nil {
		log.Printf("Could not retrived Password Hash: %v", err)
		return nil, err
	}

	// Verify the password against the hash
	accountMatch, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Printf("Error: Could not ComparePasswordAndHash successfully")
		return nil, errors.New("invalid password")
	}
	if !accountMatch {
		return nil, errors.New("invalid password")
	}

	account, err := auth.SQLite.GetAccountData(id)
	if err != nil {
		return nil, errors.New("failed to get account data")
	}

	return account, nil
}

func (auth *auth) GenerateJWT() {
	panic("Not done")
}
