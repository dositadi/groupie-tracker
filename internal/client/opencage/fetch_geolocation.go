package opencage

import (
	"net/url"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/jsonlog"
	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

type OpenCage struct {
	key    string
	logger jsonlog.Logger
}

func New(key string, logger jsonlog.Logger) OpenCage {
	return OpenCage{
		key:    key,
		logger: logger,
	}
}

type Geolocation struct {
	Name string
	Lat  string
	Lng  string
}

type raw struct {
	Results []struct {
		Geometry struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
}

func (o OpenCage) FetchGeolocation() (Geolocation, error) {
	url, err := url.Parse("https://api.opencagedata.com/geocode/v1/json")
	if err != nil {
		e := utils.WrapError("Url parse error", err)
		o.logger.PrintError(e.Error(), map[string]string{
			"Status": "Fetch geolocation f(n) under opencage pkg",
		})
		return Geolocation{}, e
	}
}
