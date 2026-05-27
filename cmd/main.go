package main

import (
	"fmt"
	"time"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
)

type Test struct{ Name string }

func main() {
	/* app := &app.App{}
	app.Run() */

	t := time.Now()

	a := artistapi.New()

	a.Init()

	fmt.Println(len(a.GetByName()))

	fmt.Println(a.GetByName())

	finished := time.Since(t)
	fmt.Println("Total time taken: ", finished.Seconds())
}
