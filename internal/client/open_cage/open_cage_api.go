package opencage

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/dositadi/groupie-tracker/internal/helper"
	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

var logger = jsonlog.New(os.Stdout, jsonlog.LevelInfo)

type GeoLocation struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lng"`
}

type openCageResponse struct {
	Results []map[string]any `json:"results"`
}

type OpenCage struct {
	apiKey string
}

func New(apiKey string) *OpenCage {
	return &OpenCage{
		apiKey: apiKey,
	}
}

func /* (p *OpenCage) */ FetchGeoLocations(query string) /* GeoLocation */ {
	baseUrl, err := url.Parse("https://api.opencagedata.com")
	if err != nil {
		e := helper.WrapError("Url Parse error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
	}

	baseUrl.Path = "geocode/v1/json"

	params := url.Values{}

	//POSITION_STACK_KEY
	params.Add("key", "51db440653e74f1097011b60838ede9d")
	params.Add("q", query)
	params.Add("limit", "1")

	baseUrl.RawQuery = params.Encode()
	fmt.Println(baseUrl.String())

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		e := helper.WrapError("New request error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		e := helper.WrapError("Do req error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	//geoLocations := make(map[string][]GeoLocation)

	/* json.NewDecoder(resp.Body).Decode() */

	/* err = json.NewDecoder(resp.Body).Decode(&geoLocations)
	if err != nil {
		e := helper.WrapError("JSON decode error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
	} */

	//return GeoLocation{Lat: geoLocations["data"][0].Lat, Long: geoLocations["data"][0].Long}
}
