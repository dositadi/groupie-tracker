package herokuapp

import (
	"encoding/json"
	"net/http"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

const (
	url     = "https://groupietrackers.herokuapp.com/api/artists"
	sourcef = "Fetch artists f(n) under client pkg"
	sourcep = "Populate artist's info with artist metadata f(n) under client pkg"
)

func (a *HerokuApp) fetchArtists() (map[int]artistMetaData, error) {
	resp, err := http.Get(url)
	if err != nil {
		e := utils.WrapError("Get error", err)
		return nil, e
	}

	defer resp.Body.Close()

	var artists []artistMetaData

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		e := utils.WrapError("Decode error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Source": sourcef,
		})
		return nil, e
	}

	out := make(map[int]artistMetaData)

	for _, artist := range artists {
		out[artist.Id] = artist
	}

	a.logger.PrintInfo("Artist's metadata  fetched successfully", map[string]string{
		"Source": sourcef,
	})
	return out, nil
}

func (a *HerokuApp) populateArtistInfoWithArtistMetaData(artists map[int]artistMetaData) chan ArtistInfo {
	chArtists := make(chan ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for _, artist := range artists {
		wg.Add(1)

		go func(art artistMetaData) {
			defer wg.Done()

			var artistInfo ArtistInfo

			artistInfo = populateArtistInfo(art, artistInfo)

			chArtists <- artistInfo
		}(artist)
	}

	go func() {
		wg.Wait()

		close(chArtists)
	}()

	a.logger.PrintInfo("Population of artist's info with the artist metadata done successfully", map[string]string{
		"Source": sourcep,
	})

	return chArtists
}
