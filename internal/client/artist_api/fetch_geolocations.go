package artistapi

import (
	"context"
	"html/template"
	"sync"

	opencage "github.com/dositadi/groupie-tracker/internal/client/open_cage"
	"github.com/dositadi/groupie-tracker/internal/helper"
)

func fillGeolocationsFromOpenCage(ctx context.Context, chArtists chan *ArtistInfo, chError chan error) chan *ArtistInfo {
	out := make(chan *ArtistInfo, 5)
	outerWg := new(sync.WaitGroup)

	for artist := range chArtists {
		current := artist
		outerWg.Add(1)

		go func(art ArtistInfo) {
			defer outerWg.Done()
			innerWg := new(sync.WaitGroup)
			var locationsGeo []opencage.GeoLocation
			mu := new(sync.RWMutex)

			for _, location := range art.Locations {
				currentLocation := location
				innerWg.Add(1)

				go func(city string) {
					defer innerWg.Done()

					cleanedCity := cleanCity(city)

					geoLocation, err := opencage.FetchGeoLocations(cleanedCity)
					if err != nil {
						chError <- err
						return
					}

					// stop go routine if an error occurs
					if err = ctx.Err(); err != nil {
						return
					}

					// Lock the map resource to avoid race conflict
					mu.Lock()
					locationsGeo = append(locationsGeo, geoLocation)
					mu.Unlock()
					// Unlock after write is done.
				}(currentLocation)
			}

			innerWg.Wait()

			// Also stop parent go routine if an error occurs from child goroutine
			if err := ctx.Err(); err != nil {
				return
			}

			jsObject := helper.Marshal(locationsGeo)

			art.GeoLocations = template.JS(jsObject)

			out <- &art
		}(*current)
	}

	go func() {
		outerWg.Wait()
		close(out)
	}()

	logger.PrintInfo("Geolocations fetch successful", map[string]string{
		"Source": "Fill geolocations f(n) under artistapi package.",
	})

	return out
}
