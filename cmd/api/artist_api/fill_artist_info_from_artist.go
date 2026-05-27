package artistapi

import (
	"sync"
)

func (a *ArtistInfo) FillArtistsInfoFromArtists(ch chan ArtistInfo, artists []Artist) chan ArtistInfo {
	temp := make(chan ArtistInfo, 5)

	var wg *sync.WaitGroup = &sync.WaitGroup{}

	for _, artist := range artists {
		wg.Add(1)

		go func(a Artist) {
			defer wg.Done()

			var artistInfo *ArtistInfo = new(ArtistInfo)

			artistInfo = PopulateArtistInfo(a, artistInfo)

			temp <- *artistInfo
		}(artist)
	}

	go func() {
		wg.Wait()
		close(temp)
	}()

	for artistInfo := range temp {
		ch <- artistInfo
	}

	logger.PrintInfo("Filled in artists in artist's info successfully", map[string]string{
		"Source": "Fill artists info from artists f(n) in artistapi pkg",
	})

	return ch
}
