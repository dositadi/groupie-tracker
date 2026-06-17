package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func (s *Server) Start() {
	s.logger.PrintInfo("api server starting", map[string]string{
		"addr": s.addr,
	})

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)

	server := http.Server{
		Addr:         s.addr,
		Handler:      s.mux,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		IdleTimeout:  1 * time.Minute,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.logger.PrintFatal("Server failed"+err.Error(), map[string]string{
				"Context": "server.Start()",
			})
			os.Exit(1)
		}
	}()

	<-chSignal

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		s.logger.PrintFatal("Server Shutdown forcefully"+err.Error(), map[string]string{
			"Context": "server.Start()",
		})
		os.Exit(1)
	}
	s.logger.PrintInfo("Server shutdown gracefully", nil)
}
