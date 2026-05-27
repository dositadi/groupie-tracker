package artistapi

import (
	"os"
	"sync"
)

func (a *ArtistInfo) fillArtistsInfoFromArtists(artists map[int]artist) chan ArtistInfo {
	temp := make(chan ArtistInfo, len(artists))

	if artists == nil {
		logger.PrintFatal("Recieved a nil paramter [artists]", map[string]string{
			"Source": "Fill artists info from artists f(n) in artistapi pkg",
		})
		os.Exit(1)
		return nil
	}

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for _, art := range artists {
		wg.Add(1)

		go func(a artist) {
			defer wg.Done()

			var artInfo *ArtistInfo = new(ArtistInfo)

			artInfo = populateArtistInfo[artist](a, artInfo)

			temp <- *artInfo
		}(art)
	}

	go func() {
		wg.Wait()
		close(temp)
	}()

	logger.PrintInfo("Filled in artists into artist's info successfully", map[string]string{
		"Source": "Fill artists info from artists f(n) in artistapi pkg",
	})

	return temp
}
