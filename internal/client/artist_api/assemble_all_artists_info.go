package artistapi

import (
	"context"
	"os"
	"time"
)

func (a *ArtistInfo) assembleArtistInfoAsMap(chArtistsInfo chan *ArtistInfo) (map[int]ArtistInfo, map[int]ArtistInfo, map[string]ArtistInfo, map[string]ArtistInfo) {
	byId := make(map[int]ArtistInfo)
	byCreationDate := make(map[int]ArtistInfo)
	byName := make(map[string]ArtistInfo)
	byFirstAlbum := make(map[string]ArtistInfo)

	for artistInfo := range chArtistsInfo {
		byId[artistInfo.id] = *artistInfo
		byCreationDate[artistInfo.creationDate] = *artistInfo
		byName[artistInfo.name] = *artistInfo
		byFirstAlbum[artistInfo.firstAlbum] = *artistInfo
	}
	return byId, byCreationDate, byName, byFirstAlbum
}

func (a *ArtistInfo) mapArtistsInfo() (map[int]ArtistInfo, map[int]ArtistInfo, map[string]ArtistInfo, map[string]ArtistInfo) {
	// Using the pipeline routine pattern to generate the artist's info
	arts, err := a.fetchArtists()
	if err != nil {
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	chError := make(chan error)

	filledArtists := a.fillArtistsInfoFromArtists(arts)
	filledLocations := a.fillArtistInfoFromLocation(ctx, filledArtists, chError, arts)
	filledDates := a.fillArtistInfoFromDate(ctx, filledLocations, chError, arts)
	filledRelations := a.fillArtistInfoFromRelations(ctx, filledDates, chError, arts)

	select {
	case <-chError:
		cancel()
		time.Sleep(5 * time.Millisecond)
		os.Exit(1)
	default:
		return a.assembleArtistInfoAsMap(filledRelations)
	}
	return a.assembleArtistInfoAsMap(filledRelations)
}
