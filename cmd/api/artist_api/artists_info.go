package artistapi

import (
	"context"
	"os"
)

func (a *ArtistInfo) InitArtistsInfo() (map[int]ArtistInfo, map[int]ArtistInfo, map[string]ArtistInfo, map[string]ArtistInfo) {
	// Using the pipeline routine pattern to generate the artist's info
	artists, err := a.fetchArtists()
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	chError := make(chan error)

	filledArtists := a.fillArtistsInfoFromArtists(artists)
	filledLocations := a.fillArtistInfoFromLocation(ctx, filledArtists, chError, artists)
	filledDates := a.fillArtistInfoFromDate(ctx, filledLocations, chError, artists)
	filledRelations := a.fillArtistInfoFromRelations(ctx, filledDates, chError, artists)

	select {
	case <-chError:
		cancel()
		os.Exit(1)
	default:
		return a.assembleArtistInfoAsMap(filledRelations)
	}
	return a.assembleArtistInfoAsMap(filledRelations)
}
