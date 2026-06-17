package herokuapp

import (
	"context"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

const (
	sourcePR = "Populate Artist Info with relations f(n) under client pkg"
)

func (a *HerokuApp) populateArtistInfoWithRelations(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error, artists map[int]artistMetaData) chan ArtistInfo {
	out := make(chan ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for artistInfo := range chArtistInfo {
		artist := artists[artistInfo.Id]

		wg.Add(1)

		go func(artInfo ArtistInfo, art artistMetaData) {
			defer wg.Done()

			relations, err := fetchInfo[relations](art.Relations)
			if err != nil {
				e := utils.WrapError("Relations fetch error", err)
				a.logger.PrintError(e.Error(), map[string]string{
					"Source": sourcePR,
				})

				chError <- e
				return
			}

			if err = ctx.Err(); err != nil {
				e := utils.WrapError("Relations fetch (ctx) error", err)

				a.logger.PrintError(e.Error(), map[string]string{
					"Source": sourcePR,
				})
				return
			}

			artInfo = populateArtistInfo(relations, artInfo)

			out <- artInfo
		}(artistInfo, artist)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	a.logger.PrintInfo("Artist's info relations updated", map[string]string{
		"Source": sourcePR,
	})

	return out
}
