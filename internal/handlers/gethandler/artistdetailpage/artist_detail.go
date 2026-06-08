package artistdetailpage

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/handlers/gethandler/homepage"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type DetailPage struct {
	logger          jsonlog.Logger
	usermodel       homepage.UserModel
	client          artistapi.ArtistInfo
	embedded        groupietracker.Embedded
	favoriteModel   homepage.FavoriteModel
	preferencemodel homepage.PreferenceModel
}

func New(usermodel homepage.UserModel, client artistapi.ArtistInfo, logger jsonlog.Logger, embedded groupietracker.Embedded, favoriteModel homepage.FavoriteModel, preferencemodel homepage.PreferenceModel) *DetailPage {
	return &DetailPage{
		usermodel:       usermodel,
		client:          client,
		logger:          logger,
		embedded:        embedded,
		favoriteModel:   favoriteModel,
		preferencemodel: preferencemodel,
	}
}
