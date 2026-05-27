package artistapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type InfoTypes interface {
	Artist | Location | ConcertDate | Relations
}

// Writing an interace type to populate the Artist struct with all it's data.
func PopulateArtistInfo[T InfoTypes](info T, artistInfo *ArtistInfo) *ArtistInfo {
	switch v := any(info).(type) {
	case Artist:
		artistInfo.Id = v.Id
		artistInfo.Image = v.Image
		artistInfo.Name = v.Name
		artistInfo.Members = v.Members
		artistInfo.CreationDate = v.CreationDate
		artistInfo.FirstAlbum = v.FirstAlbum

	case Location:
		if artistInfo != nil {
			if artistInfo.Id != v.Id {
				return nil
			}
		} else {
			return nil
		}

		artistInfo.Locations = v.Locations

	case ConcertDate:
		if artistInfo != nil {
			if artistInfo.Id != v.Id {
				return nil
			}
		} else {
			return nil
		}

		artistInfo.ConcertDates = v.Dates

	case Relations:
		if artistInfo != nil {
			if artistInfo.Id != v.Id {
				return nil
			}
		} else {
			return nil
		}

		artistInfo.DatesLocations = v.DatesLocations

	default:
		if artistInfo != nil {
			return artistInfo
		}
		return nil
	}
	return artistInfo
}

func fetchInfo[T InfoTypes](url string) (T, error) {
	var locations T
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	resp, err := http.Get(url)
	if err != nil {
		e := fmt.Errorf("Get error: %w", err)

		logger.PrintError(e.Error(), map[string]string{
			"Source": "Fetch info f(n)",
		})
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		e := fmt.Errorf("Json Decode error: %w", err)

		logger.PrintError(e.Error(), map[string]string{
			"Source": "Fetch info f(n)",
		})
	}

	return locations, nil
}
