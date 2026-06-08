package artistdetail

import (
	"net/http"

	groupietracker "github.com/dositadi/groupie-tracker"
	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type ArtistDetail struct {
	logger         jsonlog.Logger
	responseWriter http.ResponseWriter
	embedded       groupietracker.Embedded
	client         artistapi.ArtistInfo
	request        *http.Request
}

func New(logger jsonlog.Logger, responseWriter http.ResponseWriter, embedded groupietracker.Embedded, client artistapi.ArtistInfo, request *http.Request) *ArtistDetail {
	return &ArtistDetail{
		logger:         logger,
		responseWriter: responseWriter,
		embedded:       embedded,
		client:         client,
		request:        request,
	}
}
