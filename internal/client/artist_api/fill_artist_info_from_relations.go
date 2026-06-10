package artistapi

import (
	"context"
	"fmt"
	"sync"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func fillArtistInfoFromRelations(ctx context.Context, chArtistInfo chan *ArtistInfo, chError chan error, artists map[int]artist) chan *ArtistInfo {
	temp := make(chan *ArtistInfo, len(artists))
	wg := new(sync.WaitGroup)

	for artInfo := range chArtistInfo {
		art := artists[artInfo.Id]
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

			if err := ctx.Err(); err != nil {
				e := helper.WrapError("Stopping location fetch worker routine", err)

				logger.PrintFatal(e.Error(), map[string]string{
					"Source": "Fill artist info from sub infos f(n) under artistapi pkg",
					"Worker": fmt.Sprintf("Location filler for %v with %v", artInfo, a),
				})
				return
			}

			filledArtistInfo := populateArtistInfo(relation, aInfo)

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
