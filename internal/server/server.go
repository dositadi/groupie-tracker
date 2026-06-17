package server

import (
	"net/http"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/herokuapp"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/server/handlers"
)

type Server struct {
	addr      string
	logger    *jsonlog.Logger
	mux       *http.ServeMux
	templates *TemplateEngine
}

func New(addr string, logger *jsonlog.Logger, artistClient *herokuapp.HerokuApp) *Server {
	mux := http.NewServeMux()

	// static assets
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	templates := NewTemplateEngine()
	if err := templates.Load("web/templates/*.html"); err != nil {
		logger.PrintFatal(err.Error(), nil)
	}

	artistHandlers := handlers.NewArtistHandlers(artistClient, templates)
	mux.HandleFunc("GET /artists", artistHandlers.GetArtists)
	mux.HandleFunc("GET /artists/{id}", artistHandlers.GetArtistByID)

	return &Server{
		addr:      addr,
		logger:    logger,
		mux:       mux,
		templates: templates,
	}
}

func (s *Server) Templates() *TemplateEngine {
	return s.templates
}

func (s *Server) Start() error {
	s.logger.PrintInfo("api server starting", map[string]string{
		"addr": s.addr,
	})

	//signal := make(chan os.Signal, 1)

	return http.ListenAndServe(s.addr, s.mux)
}
