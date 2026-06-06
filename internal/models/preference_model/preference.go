package preferencemodel

import (
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/jackc/pgx/v5"
)

type PreferenceModel struct {
	db     *pgx.Conn
	logger jsonlog.Logger
}

func New(db *pgx.Conn, logger jsonlog.Logger) *PreferenceModel {
	return &PreferenceModel{
		logger: logger,
		db:     db,
	}
}
