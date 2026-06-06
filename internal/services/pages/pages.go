package pages

import (
	"html/template"
	"net/http"
	"strings"
	"unicode"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
	"github.com/dositadi/groupie-tracker/internal/utils"
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
	Exists(artistId int) (bool, error)
	Get(artistId int, userId string) (data.Favorite, error)
	GetAll(userId string) ([]data.Favorite, error)
	Insert(favorite data.Favorite) error
	Update(fav data.FavoriteUpdate) error
}

type Pages struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	client         artistapi.ArtistInfo
	request        *http.Request
	favoriteModel  FavoriteModel
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client artistapi.ArtistInfo, request *http.Request, favoriteModel FavoriteModel) *Pages {
	return &Pages{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		client:         client,
		request:        request,
		favoriteModel:  favoriteModel,
	}
}

func (p *Pages) getUserId() string {
	id := p.request.Context().Value(utils.USER_ID_KEY)
	if userId, ok := id.(string); ok {
		return userId
	}
	return ""
}

func (p *Pages) homePageFunc() template.FuncMap {
	return template.FuncMap{
		"GetLocation": func(s []string) string {
			return s[0]
		},
		"CleanText": func(s string) string {
			s = strings.ReplaceAll(s, "_", " ")
			s = strings.ReplaceAll(s, "-", " ")
			s = strings.ToLower(s)
			sl := strings.Fields(s)

			for i, w := range sl {
				rn := []rune(w)
				rn[0] = unicode.ToUpper(rn[0])
				sl[i] = string(rn)
			}

			return strings.Join(sl, " ")
		},
		"UpdateFavoriteField": func(artist artistapi.ArtistInfo, favorites map[int]bool) artistapi.ArtistInfo {
			if status, ok := favorites[artist.Id]; ok {
				artist.IsFavorited = status
			}
			return artist
		},
	}
}
