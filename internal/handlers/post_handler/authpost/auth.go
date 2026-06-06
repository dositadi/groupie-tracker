package authpost

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

type PreferenceModel interface {
	Exists(userId string) (bool, error)
	Get(userId string) (data.Preference, error)
	Insert(preference data.Preference) error
	Update(preference data.PreferenceUpdate) error
}

type Auth struct {
	logger          jsonlog.Logger
	usermodel       UserModel
	preferenceModel PreferenceModel
	embedded        groupietracker.Embedded
}

func New(logger jsonlog.Logger, userModel UserModel, preferenceModel PreferenceModel, embedded groupietracker.Embedded) *Auth {
	return &Auth{
		logger:          logger,
		usermodel:       userModel,
		preferenceModel: preferenceModel,
		embedded:        embedded,
	}
}
