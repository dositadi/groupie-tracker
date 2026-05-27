package main

import (
	"fmt"
	"sync"

	artistapi "github.com/dositadi/groupie-tracker/cmd/api/artist_api"
)

type Test struct{ Name string }

func main() {
	/* app := &app.App{}
	app.Run() */

	a := artistapi.ArtistInfo{}

	artists, _ := a.FetchArtists()

	ch := make(chan artistapi.ArtistInfo, 52)
	chErr := make(chan error)
	var mu *sync.Mutex = &sync.Mutex{}

	ch = a.FillArtistsInfoFromArtists(ch, chErr, *artists, mu)

	close(ch)
	fmt.Println(len(ch))

	for v := range ch {
		fmt.Println(v)
	}

}
