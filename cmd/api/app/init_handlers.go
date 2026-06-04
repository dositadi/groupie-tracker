package app

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) fileServer() {
	a.router.Handle("/src/output.css", http.FileServerFS(a.embedded.Get()))
}

func (a *App) initHandlers() {
	a.router.Use(middleware.CleanPath)
	a.fileServer()

	// Get request routes
	a.router.Group(func(r chi.Router) {
		// Auth routes
		r.Get(utils.LOGIN.String(), a.handler.Get.Auth.LoginPageHandler)
		r.Get(utils.REGISTER.String(), a.handler.Get.Auth.SignupHandler)

		// App pages
		r.With(a.midleware.Recover).With(a.midleware.VerifyAccessToken).Get(utils.HOME.String(), a.handler.Get.Pages.HomeHandler)
	})

	// Post request routes
	a.router.Group(func(r chi.Router) {
		// Auth routes
		r.With(a.midleware.Recover).Post(utils.REGISTER.String(), a.handler.Post.Auth.RegisterHandler)
		r.With(a.midleware.Recover).Post(utils.LOGIN.String(), a.handler.Post.Auth.LoginHandler)

		
		r.With(a.midleware.Recover).With(a.midleware.VerifyAccessToken).Post(utils.FAVORITE.String(), a.handler.Post.Pages.UpdateFavoriteHandler)
	})
}
