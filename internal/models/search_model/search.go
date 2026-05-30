package searchmodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5"
)

type SearchModel struct {
	db     *pgx.Conn
	logger jsonlog.Logger
}

func New(db *pgx.Conn, logger jsonlog.Logger) *SearchModel {
	return &SearchModel{
		logger: logger,
		db:     db,
	}
}
