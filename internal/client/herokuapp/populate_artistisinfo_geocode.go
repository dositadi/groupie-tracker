package herokuapp

import (
	"context"
	"html/template"
	"strings"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

var sourceG = "herokuapp.populateArtistInfoWithGeolocations()"

func (h *HerokuApp) PopulateArtistInfoWithGeolocations(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error) chan ArtistInfo {
	out := make(chan ArtistInfo)
	outerWg := new(sync.WaitGroup)

	for artistInfo := range chArtistInfo {
		outerWg.Add(1)

		go func(artistInfo ArtistInfo) {
			defer outerWg.Done()
			innerWg := new(sync.WaitGroup)

			for _, location := range artistInfo.Locations {
				innerWg.Add(1)

				go func(location string) {
					defer innerWg.Done()
					location = cleanLocation(location)

					geolocation, err := h.opencage.FetchGeolocation(location)
					if err != nil {
						e := utils.WrapError("geolocation fetch failure from worker", err)
						h.logger.PrintError(e.Error(), map[string]string{
							"Context": sourceG,
						})
						chError <- e
					}

					if err = ctx.Err(); err != nil {
						return
					}

					geoByte := utils.MarshalObject(geolocation)

					artistInfo.Geolocations = template.JS(geoByte)
				}(location)

				innerWg.Wait()

				if err := ctx.Err(); err != nil {
					return
				}

				out <- artistInfo
			}
		}(artistInfo)
	}

	go func() {
		outerWg.Wait()
		close(out)
	}()

	return out
}

func cleanLocation(location string) string {
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.ReplaceAll(location, "-", ", ")
	return location
}
