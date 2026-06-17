package herokuapp

import (
	"context"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

const (
	sourcePa = "Populate Artist Info Locations f(n) under client"
)

func (a *ArtistInfo) populateArtistInfoLocations(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error, artists map[int]artistMetaData) chan ArtistInfo {
	out := make(chan ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for artistInfo := range chArtistInfo {
		artist := artists[artistInfo.Id]
		wg.Add(1)

		go func(artInfo ArtistInfo, art artistMetaData) {
			defer wg.Done()

			locations, err := fetchInfo[locations](art.Locations)
			if err != nil {
				e := utils.WrapError("Location fetch err", err)
				logger.PrintError(e.Error(), map[string]string{
					"Source": sourcePa,
				})

				chError <- e
			}

			if err = ctx.Err(); err != nil {
				e := utils.WrapError("Location fetch (ctx) err", err)
				logger.PrintError(e.Error(), map[string]string{
					"Source": sourcePa,
				})

				return
			}

			artInfo = populateArtistInfo(locations, artInfo)

			out <- artInfo
		}(artistInfo, artist)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	logger.PrintInfo("ArtistInfos Locations updated", map[string]string{
		"Source": sourcePa,
	})
	return out
}
