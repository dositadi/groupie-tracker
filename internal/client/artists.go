package client

import (
	"context"
	"fmt"
	"os"
	"time"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
)

var (
	logger = jsonlog.New(os.Stdout, jsonlog.INFO)
)

const (
	source = "Assemble function under client pkg"
)

var byId = make(map[int]ArtistInfo)
var byName = make(map[string]ArtistInfo)
var byCreationDate = make(map[int]ArtistInfo)
var byFirstAlbum = make(map[string]ArtistInfo)

func New() *ArtistInfo {
	return &ArtistInfo{}
}

func (a *ArtistInfo) InitClient() {
	a.assemble()
}

func (a *ArtistInfo) assemble() {
	artistMetaData, err := a.fetchArtists()
	if err != nil {
		logger.PrintError(err.Error(), map[string]string{
			"Source": source,
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chError := make(chan error)

	chArtistInfo := a.populateArtistInfoWithArtistMetaData(artistMetaData)
	chArtistInfo = a.populateArtistInfoLocations(ctx, chArtistInfo, chError, artistMetaData)
	chArtistInfo = a.populateArtistInfoWithRelations(ctx, chArtistInfo, chError, artistMetaData)
	chArtistInfo = a.populateArtistInfoWithDateLocations(ctx, chArtistInfo, chError, artistMetaData)

	select {
	case <-chError:
		time.Sleep(2 * time.Second)
		logger.PrintError("An error occured with on of the worker routines", map[string]string{
			"Source": source,
		})
		os.Exit(1)
	default:
		for artistInfo := range chArtistInfo {
			byId[artistInfo.Id] = artistInfo
			byCreationDate[artistInfo.CreationDate] = artistInfo
			byFirstAlbum[artistInfo.FirstAlbum] = artistInfo
			byName[artistInfo.Name] = artistInfo
			fmt.Println(artistInfo)
		}
		logger.PrintInfo("Artist Info fetched completely", map[string]string{
			"Source": source,
		})
	}
}

func (a *ArtistInfo) GetById() map[int]ArtistInfo {
	return byId
}

func (a *ArtistInfo) GetByName() map[string]ArtistInfo {
	return byName
}

func (a *ArtistInfo) GetByFirstAlbum() map[string]ArtistInfo {
	return byFirstAlbum
}

func (a *ArtistInfo) GetByCreationDate() map[int]ArtistInfo {
	return byCreationDate
}
