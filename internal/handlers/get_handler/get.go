package gethandler

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

type Get struct {
	logger    jsonlog.Logger
	usermodel UserModel
	client    artistapi.ArtistInfo
}

func New(usermodel UserModel, client artistapi.ArtistInfo, logger jsonlog.Logger) *Get {
	return &Get{
		usermodel: usermodel,
		client:    client,
		logger:    logger,
	}
}
