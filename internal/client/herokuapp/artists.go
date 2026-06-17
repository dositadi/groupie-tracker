package herokuapp

import (
	"context"
	"os"
	"time"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/client/opencage"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
)

var (
	logger = jsonlog.New(os.Stdout, jsonlog.INFO)
)

const (
	source = "Assemble function under client pkg"
)

var byId = make(map[int]ArtistInfo)

type HerokuApp struct {
	opencage opencage.OpenCage
}

func New(opencage opencage.OpenCage) *HerokuApp {
	return &HerokuApp{
		opencage: opencage,
	}
}

func (a *HerokuApp) InitClient() {
	a.assemble()
}

func (a *HerokuApp) assemble() {
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
		}
		logger.PrintInfo("Artist Info fetched completely", map[string]string{
			"Source": source,
		})
	}
}

func (a *HerokuApp) GetById() map[int]ArtistInfo {
	return byId
}
