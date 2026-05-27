package artistapi

import (
	"context"
	"sync"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func (a *ArtistInfo) FillArtistInfoFromRelations(ctx context.Context, chArtistInfo chan *ArtistInfo, chError chan error, artists map[int]Artist) chan *ArtistInfo {
	temp := make(chan *ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for artistInfo := range chArtistInfo {
		artist := artists[artistInfo.Id]
		wg.Add(1)

		go func(aInfo *ArtistInfo, a Artist) {
			defer wg.Done()

			relation, err := fetchInfo[Relations](a.Relations)
			if err != nil {
				e := helper.WrapError("Fetch info error", err)

				logger.PrintError(e.Error(), map[string]string{
					"Source": "Fill artist info from relations f(n) in artistapi",
				})

				chError <- e
			}

			filledArtistInfo := populateArtistInfo[Relations](relation, artistInfo)

			temp <- filledArtistInfo
		}(artistInfo, artist)
	}

	go func() {
		wg.Wait()
		close(temp)
	}()

	logger.PrintInfo("Filled in relations into artist's info successfully", map[string]string{
		"Source": "Fill artist info from relations f(n) in artistapi",
	})

	return temp
}
