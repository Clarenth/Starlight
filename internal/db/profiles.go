package db

import (
	"math/rand"
	"starlight/internal/models"
	"strconv"
	"time"
)

func (c *DB) CreateNewProfile(username string, password string) {
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
