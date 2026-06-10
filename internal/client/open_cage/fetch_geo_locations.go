package opencage

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/dositadi/groupie-tracker/internal/helper"
)

func /* (p *OpenCage) */ FetchGeoLocations(query string) (GeoLocation, error) {
	baseUrl, err := url.Parse("https://api.opencagedata.com")
	if err != nil {
		e := helper.WrapError("Url Parse error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
		return GeoLocation{}, err
	}

	baseUrl.Path = "geocode/v1/json"

	params := url.Values{}

	//Set parameter values
	params.Add("key", "51db440653e74f1097011b60838ede9d")
	params.Add("q", query)
	params.Add("limit", "1")

	baseUrl.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		e := helper.WrapError("New request error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
		return GeoLocation{}, e
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		e := helper.WrapError("Do req error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
		return GeoLocation{}, e
	}

	defer resp.Body.Close()

	var openCageResp openCageRawResponse

	err = json.NewDecoder(resp.Body).Decode(&openCageResp)
	if err != nil {
		e := helper.WrapError("JSON decode error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
		return GeoLocation{}, e
	}

	geoLocations, err := parseGeoLocation(openCageResp)
	if err != nil {
		e := helper.WrapError("Geolocation parse error", err)
		logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Fetch Geo Locations f(n) under artistapi pkg",
		})
		return GeoLocation{}, e
	}

	return geoLocations, nil
}

func parseGeoLocation(raw openCageRawResponse) (GeoLocation, error) {
	if !(len(raw.Results) > 0) {
		return GeoLocation{}, errors.New("Empty response")
	}

	return GeoLocation{
		Lat: raw.Results[0].Geometry.Lat,
		Lng: raw.Results[0].Geometry.Lng,
	}, nil
}
