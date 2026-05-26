package handlers

import (
	"net/http"

	jsonlog "github.com/dositadi/groupie-tracker.git/internal/json_log"
)

type Handler struct {
	logger jsonlog.Logger
}

func New(logger jsonlog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("An internal server error occurred."))
}
