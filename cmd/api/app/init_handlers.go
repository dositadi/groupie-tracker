package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) initHandlers() {
	a.router.Use(middleware.CleanPath)
	// Get request routes
	a.router.Group(func(r chi.Router) {
		r.Get("/artists", a.handler.Get.ArtistsHandler)
	})
}
