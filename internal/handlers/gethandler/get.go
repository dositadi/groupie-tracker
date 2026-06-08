package gethandler

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/handlers/gethandler/artistdetailpage"
	"github.com/dositadi/groupie-tracker/internal/handlers/gethandler/getauth"
	"github.com/dositadi/groupie-tracker/internal/handlers/gethandler/homepage"
	"github.com/dositadi/groupie-tracker/internal/handlers/posthandler/homepagepost"
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
	Auth       getauth.Auth
	HomePage   homepage.HomePage
	DetailPage artistdetailpage.DetailPage
}

func New(usermodel UserModel, favoriteModel homepage.FavoriteModel, preferenceModel homepagepost.PreferenceModel, client artistapi.ArtistInfo, logger jsonlog.Logger, embedded groupietracker.Embedded) *Get {
	return &Get{
		Auth:       *getauth.New(usermodel, client, logger, embedded),
		HomePage:   *homepage.New(usermodel, client, logger, embedded, favoriteModel, preferenceModel),
		DetailPage: *artistdetailpage.New(usermodel, client, logger, embedded, favoriteModel, preferenceModel),
	}
}
