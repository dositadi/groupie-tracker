package favoritemodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5"
)

type FavoriteModel struct {
	db     *pgx.Conn
	logger jsonlog.Logger
}

func New(db *pgx.Conn, logger jsonlog.Logger) *FavoriteModel {
	return &FavoriteModel{
		logger: logger,
		db:     db,
	}
}
