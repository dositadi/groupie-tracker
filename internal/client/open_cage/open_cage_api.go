package opencage

import (
	"os"

	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

var logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)

type GeoLocation struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
}

type openCageRawResponse struct {
	Results []struct {
		Geometry struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
}

type OpenCage struct {
	apiKey string
}

func New(apiKey string) *OpenCage {
	return &OpenCage{
		apiKey: apiKey,
	}
}
