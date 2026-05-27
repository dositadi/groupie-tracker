package artistapi

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/dositadi/groupie-tracker/internal/helper"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

const (
	artist = "https://groupietrackers.herokuapp.com/api/artists"
)

var logger *jsonlog.Logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)

// A generic function that fetches all the artist resource.
func (a *ArtistInfo) fetchArtists() (map[int]Artist, error) {
	response, err := http.Get(artist)
	if err != nil {
		e := helper.WrapError("Get error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch artists f(n) under artistapi package.",
		})
	}

	defer response.Body.Close()

	var artists []Artist

	err = json.NewDecoder(response.Body).Decode(&artists)
	if err != nil {
		e := helper.WrapError("JSON decode error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch artists f(n) under artistapi package.",
		})
	}

	artistsMap := make(map[int]Artist)

	for _, artist := range artists {
		artistsMap[artist.Id] = artist
	}

	logger.PrintInfo("Artists fetch successful", map[string]string{
		"Source": "Fetch artists f(n) under artistapi package.",
	})

	return artistsMap, nil
}
