package app

import "os"

type config struct {
	dsn string
}

func newConfig() config {
	c := config{}
	c.dsn = os.Getenv("POSTGRES_DB_DSN")
	return c
}
