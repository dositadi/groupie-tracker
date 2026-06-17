package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strconv"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/herokuapp"
)

type Renderer interface {
	Render(w io.Writer, name string, data any) error
}

type ArtistHandlers struct {
	client    *herokuapp.ArtistInfo
	templates Renderer
}

func NewArtistHandlers(client *herokuapp.ArtistInfo, templates Renderer) *ArtistHandlers {
	return &ArtistHandlers{
		client:    client,
		templates: templates,
	}
}

func (h *ArtistHandlers) GetArtists(w http.ResponseWriter, r *http.Request) {
	artistsByID := h.client.GetById()
	artists := make([]herokuapp.ArtistInfo, 0, len(artistsByID))

	for _, artist := range artistsByID {
		artists = append(artists, artist)
	}

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Id < artists[j].Id
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Render(w, "artists.html", artists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ArtistHandlers) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid artist id"})
		return
	}

	artist, ok := h.client.GetById()[id]
	if !ok {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "artist not found"})
		return
	}

	writeJSON(w, http.StatusOK, artist)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}
