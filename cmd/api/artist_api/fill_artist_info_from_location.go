package artistapi

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func (a *ArtistInfo) FillArtistInfoFromLocation(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error, artists map[int]Artist) chan *ArtistInfo {
	temp := make(chan *ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	if chArtistInfo == nil || artists == nil {
		logger.PrintFatal("Recieved a nil paramter [chArtist | artists]", map[string]string{
			"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
		})
		os.Exit(1)
		return nil
	}

	for artistInfo := range chArtistInfo {
		artist := artists[artistInfo.Id]

		wg.Add(1)

		go func(aInfo ArtistInfo, a Artist) {
			defer wg.Done()

			locations, err := fetchInfo[Location](a.Locations)
			if err != nil {
				e := helper.WrapError("Fetch info error", err)

				logger.PrintError(e.Error(), map[string]string{
					"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
				})

				chError <- e
			}

			if err = ctx.Err(); err != nil {
				e := helper.WrapError("Stopping location fetch worker routine", err)

				logger.PrintFatal(e.Error(), map[string]string{
					"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
					"Worker": fmt.Sprintf("Location filler for %v with %v", aInfo, a),
				})
				return
			}

			artInfo := populateArtistInfo[Location](locations, &aInfo)

			temp <- artInfo

		}(artistInfo, artist)
	}

	go func() {
		wg.Wait()
		close(temp)
	}()

	logger.PrintInfo("Filled in locations into artist's info successfully", map[string]string{
		"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
	})

	return temp
}
