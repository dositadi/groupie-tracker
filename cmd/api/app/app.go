package app

import (
	"database/sql"

	"github.com/dositadi/groupie-tracker.git/cmd/api/handlers"
	"github.com/dositadi/groupie-tracker.git/cmd/api/middleware"
	jsonlog "github.com/dositadi/groupie-tracker.git/internal/json_log"
)

type app struct {
	db        *sql.DB
	config    config
	logger    jsonlog.Logger
	handler   handlers.Handler
	midleware middleware.Middleware
}
