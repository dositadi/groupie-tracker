package app

import (
	"os"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/herokuapp"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/server"
)

type App struct {
	logger jsonlog.Logger
	client herokuapp.ArtistInfo
	server *server.Server
	config Config
}

func (a *App) initApp() {
	a.config.Init()
	a.config.Validate()
	a.client = *herokuapp.New()
	a.client.InitClient()
	a.logger = *jsonlog.New(os.Stdout, jsonlog.INFO)

	a.server = server.New(":8080", &a.logger, &a.client)
}

func (a *App) Run() {
	a.initApp()

	err := a.server.Start()
	if err != nil {
		a.logger.PrintFatal(err.Error(), map[string]string{
			"Source": "Run function under cmd/app package",
		})
	}
}
