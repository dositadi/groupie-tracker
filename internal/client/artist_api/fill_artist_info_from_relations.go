package artistapi

import (
	"context"
	"sync"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func (a *ArtistInfo) fillArtistInfoFromRelations(ctx context.Context, chArtistInfo chan *ArtistInfo, chError chan error, artists map[int]artist) chan *ArtistInfo {
	temp := make(chan *ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for artInfo := range chArtistInfo {
		art := artists[artInfo.id]
		wg.Add(1)

		go func(aInfo *ArtistInfo, a artist) {
			defer wg.Done()

			relation, err := fetchInfo[relations](a.Relations)
			if err != nil {
				e := helper.WrapError("Fetch info error", err)

				logger.PrintError(e.Error(), map[string]string{
					"Source": "Fill artist info from relations f(n) in artistapi",
				})

				chError <- e
			}

			filledArtistInfo := populateArtistInfo[relations](relation, aInfo)

			temp <- filledArtistInfo
		}(artInfo, art)
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
