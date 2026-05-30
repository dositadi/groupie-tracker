package favoritemodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5"
)

type FavoriteModel struct {
	logger jsonlog.Logger
	db     *pgx.Conn
}

func New(logger jsonlog.Logger, db *pgx.Conn) *FavoriteModel {
	return &FavoriteModel{
		logger: logger,
		db:     db,
	}
}
