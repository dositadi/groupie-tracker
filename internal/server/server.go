package server

import (
	"net/http"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/server/handlers"
)

type Server struct {
	addr   string
	logger *jsonlog.Logger
	mux    *http.ServeMux
}

func New(addr string, logger *jsonlog.Logger, artistClient *client.ArtistInfo) *Server {
	mux := http.NewServeMux()

	artistHandlers := handlers.NewArtistHandlers(artistClient)
	mux.HandleFunc("GET /artists", artistHandlers.GetArtists)
	mux.HandleFunc("GET /artists/{id}", artistHandlers.GetArtistByID)

	return &Server{
		addr:   addr,
		logger: logger,
		mux:    mux,
	}
}

func (s *Server) Start() error {
	s.logger.PrintInfo("api server starting", map[string]string{
		"addr": s.addr,
	})

	return http.ListenAndServe(s.addr, s.mux)
}
