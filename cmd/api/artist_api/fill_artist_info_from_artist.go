package artistapi

import (
	"os"
	"sync"
)

func (a *ArtistInfo) FillArtistsInfoFromArtists(artists map[int]Artist) chan ArtistInfo {
	temp := make(chan ArtistInfo, len(artists))

	if artists == nil {
		logger.PrintFatal("Recieved a nil paramter [artists]", map[string]string{
			"Source": "Fill artists info from artists f(n) in artistapi pkg",
		})
		os.Exit(1)
		return nil
	}

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for _, artist := range artists {
		wg.Add(1)

		go func(a Artist) {
			defer wg.Done()

			var artistInfo *ArtistInfo = new(ArtistInfo)

			artistInfo = populateArtistInfo[Artist](a, artistInfo)

			temp <- *artistInfo
		}(artist)
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
