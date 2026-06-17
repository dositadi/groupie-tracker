package app

import (
	"errors"
	"os"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
)

type Config struct {
	OpenCageApiKey string
	logger         jsonlog.Logger
}

func New(logger jsonlog.Logger) *Config {
	return &Config{logger: logger}
}

func (c *Config) Init() {
	temp := Config{
		OpenCageApiKey: os.Getenv("OpenCageApiKey"),
	}

	c.OpenCageApiKey = temp.OpenCageApiKey
}

func (c *Config) Validate() error {
	if c.OpenCageApiKey == "" {
		e := errors.New("OpenCageApiKey is not set.")
		c.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Validate config f(n) under app pkg",
			"Hint":   "Check the environment that the OpenCageApiKey is set",
		})
		os.Exit(1)
	}
	return nil
}
