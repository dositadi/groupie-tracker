package herokuapp

import (
	"context"
	"strings"
	"sync"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/opencage"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

func (h *HerokuApp) populateArtistInfoWithGeolocations(ctx context.Context, chArtistInfo chan ArtistInfo, chError chan error, artists map[int]ArtistInfo) chan ArtistInfo {
	out := make(chan ArtistInfo)
	outerWg := new(sync.WaitGroup)

	for artistInfo := range chArtistInfo {
		outerWg.Add(1)

		go func(artistInfo ArtistInfo) {
			defer outerWg.Done()
			innerWg := new(sync.WaitGroup)

			for _, location := range artistInfo.Locations {
				innerWg.Add(1)

				go func (location string)  {
					defer innerWg.Done()
					location = cleanLocation(location)

					geolocation, err := h.opencage.FetchGeolocation(location)
					if err != nil {
						e :=  utils.WrapError("geolocation fetch failure from worker", err)
						
					}
					
				}(location)

			}
		}(artistInfo)
	}
}

func cleanLocation(location string) string {
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.ReplaceAll(location, "-", ", ")
	return location
}
