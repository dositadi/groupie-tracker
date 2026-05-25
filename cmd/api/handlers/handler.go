package handlers

import "net/http"

type Handler struct{}

func (h *Handler) ServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("An internal server error occurred."))
}
