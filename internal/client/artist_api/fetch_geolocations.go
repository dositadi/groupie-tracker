package artistapi

import (
	"context"
	"sync"

	opencage "github.com/dositadi/groupie-tracker/internal/client/open_cage"
)

func (a *ArtistInfo) FillGeolocationsFromOpenCage(ctx context.Context, chArtists chan ArtistInfo, chError chan error) {
	out := make(chan ArtistInfo)
	wg := new(sync.WaitGroup)

	for artist := range chArtists {
		current := artist
		wg.Add(1)

		go func (art ArtistInfo)  {
			defer wg.Done()
			goWg := new(sync.WaitGroup)
			locationsGeo := make(map[string]opencage.GeoLocation)

			for _,location := range art.Locations {
				goWg.Add(1)

				city := cleanCity(location)

				geoLocation := opencage.FetchGeoLocations(city)
			}
		}(current)
	}
}
