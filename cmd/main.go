package main

import (
	"github.com/dositadi/groupie-tracker/cmd/api/app"
)

type Test struct{ Name string }

func main() {
	app := &app.App{}
	app.Run()
}
