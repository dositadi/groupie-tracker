package posthandler

import (
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

type Post struct {
	logger    jsonlog.Logger
	userModel UserModel
	client    artistapi.ArtistInfo
}

func New(userModel UserModel, client artistapi.ArtistInfo, logger jsonlog.Logger) *Post {
	return &Post{
		userModel: userModel,
		logger:    logger,
		client:    client,
	}
}
