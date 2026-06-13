package ticketpage

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type Filter string
type Sort string
type Favorite string

const (
	// Filters
	FILTER_BY_ID            Filter = "ID"
	FILTER_BY_NAME          Filter = "NAME"
	FILTER_BY_CREATION_DATE Filter = "CREATION DATE"
	FILTER_BY_FIRST_ALBUM   Filter = "FIRST ALBUM"

	// Sort orders
	ASCENDING_ORDER  Sort = "ASC"
	DESCENDING_ORDER Sort = "DESC"

	// Favorite
	FAVORITED     Favorite = "true"
	NOT_FAVORITED Favorite = "false"
)

type FavoriteModel interface {
	DeleteAll(userId string) error
	Delete(userId string, artistId string) error
	Exists(artistId int, userId string) (bool, error)
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

type TicketPage struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	client         artistapi.ArtistInfo
	request        *http.Request /*
		favoriteModel   FavoriteModel
		preferenceModel PreferenceModel */
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client artistapi.ArtistInfo, request *http.Request /* favoriteModel FavoriteModel, preferenceModel PreferenceModel */) *TicketPage {
	return &TicketPage{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		client:         client,
		request:        request, /*
			favoriteModel:   favoriteModel,
			preferenceModel: preferenceModel, */
	}
}
