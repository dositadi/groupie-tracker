package herokuapp

import (
	"context"
	"html/template"
	"strings"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/opencage"
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
			geoLocations := make([]opencage.Geolocation, 0, len(artistInfo.Locations))
			mu := new(sync.Mutex)

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
						select {
						case chError <- e:
						case <-ctx.Done():
						default:
						}
						return
					}

					if err = ctx.Err(); err != nil {
						return
					}

					mu.Lock()
					geoLocations = append(geoLocations, geolocation)
					mu.Unlock()
				}(location)
			}

			innerWg.Wait()

			if err := ctx.Err(); err != nil {
				return
			}

			geoLocationsByte := utils.MarshalObject(geoLocations)

			artistInfo.Geolocations = template.JS(geoLocationsByte)

			select {
			case out <- artistInfo:
			case <-ctx.Done():
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
