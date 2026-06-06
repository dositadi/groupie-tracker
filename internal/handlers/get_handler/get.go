package gethandler

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/handlers/get_handler/apppages"
	"github.com/dositadi/groupie-tracker/internal/handlers/get_handler/getauth"
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

type Get struct {
	Auth  getauth.Auth
	Pages apppages.Pages
}

func New(usermodel UserModel, favoriteModel apppages.FavoriteModel, client artistapi.ArtistInfo, logger jsonlog.Logger, embedded groupietracker.Embedded) *Get {
	return &Get{
		Auth:  *getauth.New(usermodel, client, logger, embedded),
		Pages: *apppages.New(usermodel, client, logger, embedded, favoriteModel),
	}
}
