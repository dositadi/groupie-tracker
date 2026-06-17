package app

import (
	"os"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/herokuapp"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/opencage"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/server"
)

type App struct {
	logger   jsonlog.Logger
	opencage opencage.OpenCage
	client   herokuapp.HerokuApp
	server   *server.Server
	config   Config
}

func (a *App) initApp() {
	a.config.Init()
	a.config.Validate()
	a.logger = *jsonlog.New(os.Stdout, jsonlog.INFO)
	a.opencage = opencage.New(a.config.OpenCageApiKey, a.logger)
	a.client = *herokuapp.New(a.opencage, a.logger)
	a.client.InitClient()

	a.server = server.New(":8080", &a.logger, a.client.GetById())
}

func (a *App) Run() {
	a.initApp()

	a.server.Start()
}
