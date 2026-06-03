package app

import (
	"os"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/handlers"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/middlewares"
	"github.com/dositadi/groupie-tracker/internal/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type App struct {
	db        *pgx.Conn
	config    config
	logger    *jsonlog.Logger
	handler   *handlers.Handler
	midleware *middlewares.Middleware
	client    *artistapi.ArtistInfo
	router    *chi.Mux
	embedded  groupietracker.Embedded
}

func (a *App) init() {
	a.embedded = *groupietracker.New()
	a.router = chi.NewRouter()
	a.client = artistapi.New()
	a.client.Init()
	a.config = newConfig()
	a.logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	a.initDB()
	models := models.New(a.db, *a.logger)
	a.handler = handlers.New(*a.logger, &models.UserModel, *a.client, a.embedded)
	a.midleware = middlewares.New(*a.handler, *a.logger)
	a.initHandlers()
}

func (a *App) Run() {
	a.init()
	a.startServer()
}
