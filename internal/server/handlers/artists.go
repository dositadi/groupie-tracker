package handlers

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/herokuapp"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
)

type Renderer interface {
	Render(w io.Writer, name string, data any) error
}

type ArtistHandlers struct {
	client    map[int]herokuapp.ArtistInfo
	templates Renderer
	logger    jsonlog.Logger
}

func NewArtistHandlers(client map[int]herokuapp.ArtistInfo, templates Renderer, logger jsonlog.Logger) *ArtistHandlers {
	return &ArtistHandlers{
		client:    client,
		templates: templates,
		logger:    logger,
	}
}

// ArtistDetail merges the cached artist with live relations data for the detail page.
type ArtistDetail struct {
	herokuapp.ArtistInfo
	DatesLocations map[string][]string
}

func (h *ArtistHandlers) GetArtists(w http.ResponseWriter, r *http.Request) {
	artists := make([]herokuapp.ArtistInfo, 0, len(h.client))

	for _, artist := range h.client {
		artists = append(artists, artist)
	}

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Id < artists[j].Id
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Render(w, "artists.html", artists); err != nil {
		h.logger.PrintError(err.Error(), map[string]string{
			"Context": "handlers.GetArtists()",
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ArtistHandlers) GetArtistByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	artist, ok := h.client[id]
	if !ok {
		http.NotFound(w, r)
		return
	}
	fmt.Println(artist)

	data := struct {
		ArtistInfo herokuapp.ArtistInfo
	}{
		ArtistInfo: artist,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Render(w, "artistdetails.html", data); err != nil {
		h.logger.PrintError(err.Error(), map[string]string{
			"Context": "handlers.GetArtistById()",
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
