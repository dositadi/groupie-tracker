package artistapi

type ArtistInfo struct {
	Id             int
	Image          string
	Name           string
	Members        []string
	CreationDate   int
	FirstAlbum     string
	Locations      []string
	ConcertDates   []string
	DatesLocations map[string][]string
}

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
