package main

import (
	"fmt"
	"os"

	"github.com/dositadi/groupie-tracker.git/internal/helper"
	jsonlog "github.com/dositadi/groupie-tracker.git/internal/json_log"
)

type Test struct{ Name string }

func main() {
	log := jsonlog.New(os.Stdout, 0)

	log.PrintError("Failed to register user", map[string]string{"name": "Desmond", "role": "Admin"})

	fmt.Println(string(helper.Marshal(Test{Name: "Samuel"})))
}
