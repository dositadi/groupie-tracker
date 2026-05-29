package usermodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5"
)

type UserModel struct {
	db     *pgx.Conn
	logger *jsonlog.Logger
}

const (
	timeOut = 5
)

func New(db *pgx.Conn, logger *jsonlog.Logger) *UserModel {
	return &UserModel{
		db:     db,
		logger: logger,
	}
}
