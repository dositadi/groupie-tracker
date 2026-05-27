package app

import (
	"os"

	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/middleware"
	"github.com/jackc/pgx/v5"
)

type App struct {
	db        *pgx.Conn
	config    config
	logger    *jsonlog.Logger
	handler   *handlers.Handler
	midleware *middleware.Middleware
}

func (a *App) init() {
	a.config = newConfig()
	a.logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	a.handler = handlers.New(*a.logger)
	a.midleware = middleware.New(*a.handler, *a.logger)
	a.initDB()
}

func (a *App) Run() {
	a.init()
}
