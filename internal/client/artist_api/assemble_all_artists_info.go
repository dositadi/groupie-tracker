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
	arts, err := a.fetchArtists()
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chError := make(chan error)

	filledArtists := a.fillArtistsInfoFromArtists(arts)
	chArtistInfo := a.fillArtistInfoFromLocation(ctx, filledArtists, chError, arts)
	chArtistInfo = a.fillArtistInfoFromDate(ctx, chArtistInfo, chError, arts)
	chArtistInfo = a.fillArtistInfoFromRelations(ctx, chArtistInfo, chError, arts)

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
