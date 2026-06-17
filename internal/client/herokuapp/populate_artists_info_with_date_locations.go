package herokuapp

import (
	"context"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

const (
	sourcePD = "Populate artist info with data locations f(n) under client"
)

func (a *HerokuApp) populateArtistInfoWithDateLocations(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error, artists map[int]artistMetaData) chan ArtistInfo {
	out := make(chan ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for artistInfo := range chArtistInfo {
		artist := artists[artistInfo.Id]

		wg.Add(1)

		go func(artInfo ArtistInfo, art artistMetaData) {
			defer wg.Done()

			concertDates, err := fetchInfo[concertDates](art.ConcertDates)
			if err != nil {
				e := utils.WrapError("Concert dates fetch error", err)
				a.logger.PrintError(e.Error(), map[string]string{
					"Source": sourcePD,
				})
				chError <- e
			}

			if err = ctx.Err(); err != nil {
				e := utils.WrapError("Concert dates fetch (ctx) error", err)

				a.logger.PrintError(e.Error(), map[string]string{
					"Source": sourcePD,
				})

				return
			}

			artInfo = populateArtistInfo(concertDates, artInfo)

			out <- artInfo
		}(artistInfo, artist)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	a.logger.PrintInfo("Artist's info date locations updated", map[string]string{
		"Source": sourcePD,
	})

	return out
}
