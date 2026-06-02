package auth

import (
	groupietracker "github.com/dositadi/groupie-tracker"
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

type Auth struct {
	logger    jsonlog.Logger
	usermodel UserModel
	embedded  groupietracker.Embedded
}

func New(logger jsonlog.Logger, userModel UserModel, embedded groupietracker.Embedded) *Auth {
	return &Auth{
		logger:    logger,
		usermodel: userModel,
		embedded:  embedded,
	}
}
