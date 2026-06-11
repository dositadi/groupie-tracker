package artistapi

import (
	"context"
	"os"
	"time"
)

var (
	byId           = make(map[int]ArtistInfo)
	byCreationDate = make(map[int]ArtistInfo)
	byName         = make(map[string]ArtistInfo)
	byFirstAlbum   = make(map[string]ArtistInfo)
)

func (a *ArtistInfo) mapArtistsInfo() {
	// Using the pipeline routine pattern to generate the artist's info
	arts, err := fetchArtists()
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chError := make(chan error)

	filledArtists := fillArtistsInfoFromArtists(arts)
	chArtistInfo := fillArtistInfoFromLocation(ctx, filledArtists, chError, arts)
	chArtistInfo = fillArtistInfoFromDate(ctx, chArtistInfo, chError, arts)
	chArtistInfo = fillArtistInfoFromRelations(ctx, chArtistInfo, chError, arts)
	//chArtistInfo = fillGeolocationsFromOpenCage(ctx, chArtistInfo, chError)

	select {
	case <-chError:
		time.Sleep(5 * time.Millisecond)
		os.Exit(1)
	default:
		for artistInfo := range chArtistInfo {
			byId[artistInfo.Id] = *artistInfo
			byCreationDate[artistInfo.CreationDate] = *artistInfo
			byName[artistInfo.Name] = *artistInfo
			byFirstAlbum[artistInfo.FirstAlbum] = *artistInfo
		}
	}
}
