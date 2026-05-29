package data

import "time"

type User struct {
	Id             string
	Username       string
	Email          string
	HashedPassword string
	Version        int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UpdateUser struct {
	Username       *string
	Email          *string
	HashedPassword *string
}
