package app

import (
	"database/sql"
	"os"

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

func (a *app) init() {
	a.config = newConfig()
	a.logger = *jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	a.handler = *handlers.New(a.logger)
	a.midleware = *middleware.New(a.handler, a.logger)
}
