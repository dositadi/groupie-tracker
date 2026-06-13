package ticketpage

import (
	"net/http"
	"strconv"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

/* type UserModel interface {
	Delete(id string) error
	GetWithID(id string) (data.User, error)
	GetWithEmail(email string) (data.User, error)
	Insert(user data.User) error
	Update(id string, info data.UpdateUser) error
	EmailExists(email string) (bool, error)
	IDExists(id string) (bool, error)
} */
/*
type PreferenceModel interface {
	Exists(userId string) (bool, error)
	Get(userId string) (data.Preference, error)
	Insert(preference data.Preference) error
	Update(preference data.PreferenceUpdate) error
} */

/*
	 type FavoriteModel interface {
		DeleteAll(userId string) error
		Delete(userId string, artistId string) error
		Exists(artistId int, userId string) (bool, error)
		Get(artistId int, userId string) (data.Favorite, error)
		GetAll(userId string) ([]data.Favorite, error)
		Insert(favorite data.Favorite) error
		Update(fav data.FavoriteUpdate) error
	}
*/
type TicketPage struct {
	logger   jsonlog.Logger
	embedded groupietracker.Embedded
	client   artistapi.ArtistInfo
}

func New(logger jsonlog.Logger, embedded groupietracker.Embedded, client artistapi.ArtistInfo) *TicketPage {
	return &TicketPage{
		logger:   logger,
		embedded: embedded,
		client:   client,
	}
}

func (t *TicketPage) getUserId(r *http.Request) data.User {
	val := r.Context().Value(utils.USER_ID_KEY)
	var user data.User
	if v, ok := val.(data.User); ok {
		user = v
	}
	return user
}

func (t *TicketPage) atoi(s string, source string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		t.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": source,
		})
		panic("Not a valid number")
	}
	return out
}
