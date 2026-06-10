package artistapi

import (
	"context"
	"sync"

	opencage "github.com/dositadi/groupie-tracker/internal/client/open_cage"
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
			locationsGeo := make(map[string]opencage.GeoLocation)
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
					locationsGeo[city] = geoLocation
					mu.Unlock()
					// Unlock after write is done.
				}(currentLocation)
			}

			innerWg.Wait()

			// Also stop parent go routine if an error occurs from child goroutine
			if err := ctx.Err(); err != nil {
				return
			}
			art.GeoLocations = locationsGeo

			out <- &art
		}(*current)
	}

	go func() {
		outerWg.Wait()
		close(out)
	}()

	return out
}
