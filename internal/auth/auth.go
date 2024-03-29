package auth

import (
	"context"
	"fmt"
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
	/*
			hashedPassword, err := argon2id.CreateHash(accountData.Password, argon2Params)
		if err != nil {
			log.Panicf("Unable to encrypt accountData password from new Signup with given email: %v\n", accountData.Email)
			return apperrors.NewInternal()
		}
		accountData.Password = hashedPassword //The hashed password is saved to the DB instead of the plain text. User Signin will compare the client's input password to the stored hash and validity
	*/

	newAccount := models.Account{
		ID:        uuid.New(),
		Username:  username,
		Password:  password,
		CreatedAt: time.UTC.String(),
		UpdatedAt: time.UTC.String(),
	}
	// fmt.Sprintln(newAccount)
	ctx := auth.ctx
	result, err := auth.SQLite.CreateAccount(ctx, &newAccount)
	if err != nil {
		panic("TODO")
	}
	return result, nil
}

func (auth *auth) Login(username string, password string) (*models.Account, error) {
	if username == "" || password == "" {
		return nil, fmt.Errorf("username and password cannot be empty")
	}

	// implement Argon2 hash and salt checking here
	// this for ensuring security
	/*

		// Verify the password against the hash
		accountMatch, err := argon2id.ComparePasswordAndHash(payload.Password, account.Password)
		if err != nil {
			log.Printf("Error: Could not ComparePasswordAndHash successfully")
			return nil, apperrors.NewInternal()
		}
		if !accountMatch {
			return nil, apperrors.NewAuthorization("Debug 2 Invalid email or password boolean")
		}

	*/

	account, err := auth.SQLite.GetAccount(auth.ctx, username, password)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (auth *auth) GenerateJWT() {
	panic("Not done")
}
