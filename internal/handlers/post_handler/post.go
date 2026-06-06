package posthandler

import (
	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/handlers/post_handler/authpost"
	"github.com/dositadi/groupie-tracker/internal/handlers/post_handler/pages"
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

type Post struct {
	logger    jsonlog.Logger
	userModel UserModel
	client    artistapi.ArtistInfo
	embedded  groupietracker.Embedded
	Auth      authpost.Auth
	Pages     pages.Pages
}

func New(userModel UserModel, favoriteModel pages.FavoriteModel, preferenceModel pages.PreferenceModel, client artistapi.ArtistInfo, logger jsonlog.Logger, embedded groupietracker.Embedded) *Post {
	return &Post{
		userModel: userModel,
		logger:    logger,
		client:    client,
		embedded:  embedded,
		Auth:      *authpost.New(logger, userModel, preferenceModel,embedded),
		Pages:     *pages.New(logger, userModel, favoriteModel, preferenceModel, embedded, client),
	}
}
