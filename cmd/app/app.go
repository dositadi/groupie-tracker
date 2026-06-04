package app

import (
	"os"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
)

type App struct {
	logger jsonlog.Logger
	client client.ArtistInfo
}

func (a *App) initApp() {
	a.client = *client.New()
	a.client.InitClient()
	a.logger = *jsonlog.New(os.Stdout, jsonlog.INFO)
}

func (a *App) Run() {
	a.initApp()
}
