package models

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	usermodel "github.com/dositadi/groupie-tracker/internal/models/user_model"
	"github.com/jackc/pgx/v5"
)

type Models struct {
	UserModel usermodel.UserModel
}

func New(db *pgx.Conn, logger *jsonlog.Logger) *Models {
	return &Models{
		UserModel: *usermodel.New(db, logger),
	}
}
