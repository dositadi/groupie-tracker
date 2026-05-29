package gethandler

import (
	"encoding/json"
	"net/http"
)

func (g *Get) ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	arts := g.client.GetByName()
	artists, err := json.Marshal(arts)
	if err != nil {
		http.Error(w, "Json Decode error", http.StatusInternalServerError)
		return
	}

	//fmt.Println(g.client.GetByName())

	w.WriteHeader(http.StatusOK)

	w.Write(artists)
}
