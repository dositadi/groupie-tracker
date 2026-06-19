package opencage

import (
	"encoding/json"
	"errors"
	"net/http"
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
	Lat  float64
	Lng  float64
}

type raw struct {
	Results []struct {
		Geometry struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
}


func (o OpenCage) FetchGeolocation(query string) (Geolocation, error) {
	path, err := url.Parse("https://api.opencagedata.com/geocode/v1/json")
	if err != nil {
		e := utils.WrapError("Url parse error", err)
		o.logger.PrintError(e.Error(), map[string]string{
			"Status": "Fetch geolocation f(n) under opencage pkg",
		})
		return Geolocation{}, e
	}

	params := url.Values{}
	params.Add("q", query)
	params.Add("key", "a45e2bfd61d04e13b6504d106de3db70")
	params.Add("limit", "1")

	path.RawQuery = params.Encode()

	req, err := http.NewRequest(http.MethodGet, path.String(), nil)
	if err != nil {
		e := utils.WrapError("http new request error", err)
		o.logger.PrintError(e.Error(), map[string]string{
			"Status": "Fetch geolocation f(n) under opencage pkg",
		})
		return Geolocation{}, e
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		e := utils.WrapError("http client request do error", err)
		o.logger.PrintError(e.Error(), map[string]string{
			"Status": "Fetch geolocation f(n) under opencage pkg",
		})
		return Geolocation{}, e
	}
	defer resp.Body.Close()

	var rawGeolocation raw

	err = json.NewDecoder(resp.Body).Decode(&rawGeolocation)
	if err != nil {
		e := utils.WrapError("response body decode error", err)
		o.logger.PrintError(e.Error(), map[string]string{
			"Status": "Fetch geolocation f(n) under opencage pkg",
		})
		return Geolocation{}, e
	}

	geolocation, err := o.flattenRawGeolocation(query, rawGeolocation)
	if err != nil {
		e := utils.WrapError("flatten response error", err)
		o.logger.PrintError(e.Error(), map[string]string{
			"Status": "Fetch geolocation f(n) under opencage pkg",
		})
		return Geolocation{}, e
	}
	return geolocation, nil
}

func (o OpenCage) flattenRawGeolocation(city string, raw raw) (Geolocation, error) {
	if len(raw.Results) < 1 {
		return Geolocation{}, errors.New("Empty geolocation result")
	}

	return Geolocation{
		Name: city,
		Lat:  raw.Results[0].Geometry.Lat,
		Lng:  raw.Results[0].Geometry.Lng,
	}, nil
}
