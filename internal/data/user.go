package data

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id             string
	Username       string
	Email          string
	HashedPassword []byte
	Version        int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UpdateUser struct {
	Username       *string
	Email          *string
	HashedPassword []byte
}

type ActiveUser struct {
	Id       string
	Username string
	Email    string
	jwt.RegisteredClaims
}
