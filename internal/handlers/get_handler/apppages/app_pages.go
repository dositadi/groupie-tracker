package apppages

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
}

type FavoriteModel interface {
	DeleteAll(userId string) error
	Delete(userId string, artistId string) error
	Exists(artistId int) (bool, error)
	Get(artistId int, userId string) (data.Favorite, error)
	GetAll(userId string) ([]data.Favorite, error)
	Insert(favorite data.Favorite) error
	Update(fav data.FavoriteUpdate) error
}

type PreferenceModel interface {
	Exists(userId string) (bool, error)
	Get(userId string) (data.Preference, error)
	Insert(preference data.Preference) error
	Update(preference data.PreferenceUpdate) error
}

type Pages struct {
	logger          jsonlog.Logger
	usermodel       UserModel
	client          artistapi.ArtistInfo
	embedded        groupietracker.Embedded
	favoriteModel   FavoriteModel
	preferencemodel PreferenceModel
}

func New(usermodel UserModel, client artistapi.ArtistInfo, logger jsonlog.Logger, embedded groupietracker.Embedded, favoriteModel FavoriteModel, preferencemodel PreferenceModel) *Pages {
	return &Pages{
		usermodel:       usermodel,
		client:          client,
		logger:          logger,
		embedded:        embedded,
		favoriteModel:   favoriteModel,
		preferencemodel: preferencemodel,
	}
}
