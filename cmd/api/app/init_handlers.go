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
	a.router.Use(a.midleware.Recover)
	a.fileServer()

	// Get request routes
	a.router.Group(func(r chi.Router) {
		// Auth routes
		r.Get(utils.LOGIN.String(), a.handler.Get.Auth.LoginPageHandler)
		r.Get(utils.REGISTER.String(), a.handler.Get.Auth.SignupHandler)

		// App pages
		r.With(a.midleware.VerifyAccessToken).Get(utils.HOME.String(), a.handler.Get.HomePage.HomeHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.ARTIST_SEARCH.String(), a.handler.Get.HomePage.SearchHandler)
		r.With(a.midleware.VerifyAccessToken).Get(utils.ARTIST_DETAILS.String()+"/{id}", a.handler.Get.DetailPage.DetailPageHandler)
	})

	// Post request routes
	a.router.Group(func(r chi.Router) {
		// Auth routes
		r.Post(utils.REGISTER.String(), a.handler.Post.Auth.RegisterHandler)
		r.Post(utils.LOGIN.String(), a.handler.Post.Auth.LoginHandler)
		r.With(a.midleware.VerifyAccessToken).Post(utils.FILTER_SORT_ROUTE.String(), a.handler.Post.HomePage.FilterSortHandler)

		// App post request
		r.With(a.midleware.VerifyAccessToken).Post(utils.FAVORITE.String(), a.handler.Post.HomePage.UpdateFavoriteHandler)
	})
}
