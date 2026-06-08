package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client"
)

type ArtistHandlers struct {
	client *client.ArtistInfo
}

func NewArtistHandlers(client *client.ArtistInfo) *ArtistHandlers {
	return &ArtistHandlers{client: client}
}

func (h *ArtistHandlers) GetArtists(w http.ResponseWriter, r *http.Request) {
	artistsByID := h.client.GetById()
	artists := make([]client.ArtistInfo, 0, len(artistsByID))

	for _, artist := range artistsByID {
		artists = append(artists, artist)
	}

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Id < artists[j].Id
	})

	writeJSON(w, http.StatusOK, artists)
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
