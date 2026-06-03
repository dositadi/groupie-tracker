package pages

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type Pages struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	client         artistapi.ArtistInfo
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client artistapi.ArtistInfo) *Pages {
	return &Pages{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		client:         client,
	}
}
