package client

import (
	"encoding/json"
	"net/http"

	"acad.learn2earn.ng/git/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceI  = "Fetch info f(n) under client pkg"
)

type infoType interface {
	locations | artistMetaData | relations | concertDates
}

func fetchInfo[T infoType](url string) (out T, err error) {
	resp, err := http.Get(url)
	if err != nil {
		e := utils.WrapError("Get error", err)
		logger.PrintError(e.Error(), map[string]string{
			"Source": sourceI,
		})
		return out, e
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&out)
	if err != nil {
		e := utils.WrapError("Decode error", err)
		logger.PrintError(e.Error(), map[string]string{
			"Source": sourceI,
		})
		return out, e
	}
	return out, nil
}

func populateArtistInfo[T infoType](info T, artistInfo ArtistInfo) ArtistInfo {
	switch v := any(info).(type) {
	case artistMetaData:
		artistInfo.Id = v.Id
		artistInfo.Image = v.Image
		artistInfo.Name = v.Name
		artistInfo.FirstAlbum = v.FirstAlbum
		artistInfo.CreationDate = v.CreationDate
		artistInfo.Members = v.Members

	case locations:
		artistInfo.Locations = v.Locations

	case concertDates:
		artistInfo.ConcertDates = v.ConcertDates

	case relations:
		artistInfo.Relations = v.Relations

	}
	return artistInfo
}
