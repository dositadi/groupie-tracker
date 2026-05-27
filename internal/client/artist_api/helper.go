package artistapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jsonlog "github.com/dositadi/groupie-tracker/internal/json_log"
)

type InfoTypes interface {
	artist | location | concertDate | relations
}

// Writing an interace type to populate the Artist struct with all it's data.
func populateArtistInfo[T InfoTypes](info T, artistInfo *ArtistInfo) *ArtistInfo {
	switch v := any(info).(type) {
	case artist:
		artistInfo.id = v.Id
		artistInfo.image = v.Image
		artistInfo.name = v.Name
		artistInfo.members = v.Members
		artistInfo.creationDate = v.CreationDate
		artistInfo.firstAlbum = v.FirstAlbum

	case location:
		if artistInfo != nil {
			if artistInfo.id == v.Id {
				artistInfo.locations = v.Locations
			}
		} else {
			return nil
		}

	case concertDate:
		if artistInfo != nil {
			if artistInfo.id == v.Id {
				artistInfo.concertDates = v.Dates
			}
		} else {
			return nil
		}

	case relations:
		if artistInfo != nil {
			if artistInfo.id == v.Id {
				artistInfo.datesLocations = v.DatesLocations
			}
		} else {
			return nil
		}

	default:
		if artistInfo != nil {
			return artistInfo
		}
		return nil
	}
	return artistInfo
}

func fetchInfo[T InfoTypes](url string) (T, error) {
	var info T
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	resp, err := http.Get(url)
	if err != nil {
		e := fmt.Errorf("Get error: %w", err)

		logger.PrintError(e.Error(), map[string]string{
			"Source": "Fetch info f(n)",
		})
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		e := fmt.Errorf("Json Decode error: %w", err)

		logger.PrintError(e.Error(), map[string]string{
			"Source": "Fetch info f(n)",
		})
	}

	return info, nil
}
