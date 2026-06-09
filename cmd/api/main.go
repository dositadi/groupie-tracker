package main

import (
	opencage "github.com/dositadi/groupie-tracker/internal/client/open_cage"
)

func main() {
	/* app := &app.App{}
	app.Run() */

	opencage.FetchGeoLocations("Recife Brazil")
}
