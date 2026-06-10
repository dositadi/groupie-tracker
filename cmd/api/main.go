package main

import (
	"fmt"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
)

func main() {
	/* app := &app.App{}
	app.Run() */

	/* opencage.FetchGeoLocations("Recife Brazil") */
	a := artistapi.New()
	a.Init()

	fmt.Println(a.GetByIdKey())
}
