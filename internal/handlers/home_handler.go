package handlers

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	json := helper.Marshal(map[string]string{"Name": "Divine"})
	w.Write(json)
}
