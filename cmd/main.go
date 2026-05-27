package main

import (
	"context"
	"fmt"
	"os"
	"time"

	artistapi "github.com/dositadi/groupie-tracker/cmd/api/artist_api"
)

type Test struct{ Name string }

func main() {
	/* app := &app.App{}
	app.Run() */

	t := time.Now()

	a := artistapi.ArtistInfo{}

	artists, _ := a.FetchArtists()

	chError := make(chan error)

	ctx, cancel := context.WithCancel(context.Background())

	artistFilled := a.FillArtistsInfoFromArtists(artists)
	locationFilled := a.FillArtistInfoFromLocation(ctx, artistFilled, chError, artists)
	dateFilled := a.FillArtistInfoFromDate(ctx, locationFilled, chError, artists)
	relationFilled := a.FillArtistInfoFromRelations(ctx, dateFilled, chError, artists)

	fmt.Println(len(relationFilled))

	for v := range relationFilled {
		fmt.Println(v)
	}

	select {
	case <-chError:
		cancel()
		time.Sleep(10 * time.Second)
		os.Exit(1)
	default:
		finished := time.Since(t)
		fmt.Println("Total time taken: ", finished.Seconds())
		return
	}
}
