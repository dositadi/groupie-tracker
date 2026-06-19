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
	client    map[int]herokuapp.ArtistInfo
	templates Renderer
}

func NewArtistHandlers(client map[int]herokuapp.ArtistInfo, templates Renderer) *ArtistHandlers {
	return &ArtistHandlers{
		client:    client,
		templates: templates,
	}
}

// ArtistDetail merges the cached artist with live relations data for the detail page.
type ArtistDetail struct {
	herokuapp.ArtistInfo
	DatesLocations map[string][]string
}

type relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	rel, err := fetchRelations(id)
	if err != nil {
		http.Error(w, "could not load tour data", http.StatusBadGateway)
		return
	}

	detail := ArtistDetail{
		ArtistInfo:     artist,
		DatesLocations: rel.DatesLocations,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.templates.Render(w, "artistdetails.html", detail); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func fetchRelations(id int) (relation, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(id)

	resp, err := http.Get(url)
	if err != nil {
		return relation{}, err
	}
	defer resp.Body.Close()

	var rel relation
	if err := json.NewDecoder(resp.Body).Decode(&rel); err != nil {
		return relation{}, err
	}

	return rel, nil
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
