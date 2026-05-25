package main

import (
	"os"

	jsonlog "github.com/dositadi/groupie-tracker.git/internal/json_log"
)

func main() {
	log := jsonlog.New(os.Stdout, 0)

	log.PrintError("Failed to register user", map[string]string{"name": "Desmond", "role": "Admin"})
}
