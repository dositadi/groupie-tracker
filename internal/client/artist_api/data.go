package artistapi

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`    // Further get request
	ConcertDates string   `json:"concertDates"` // Further get request
	Relations    string   `json:"relations"`    // Further get request
}

type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"` // Further get request
}

type concertDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Flattened structure
type ArtistInfo struct {
	Id             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creation_date"`
	FirstAlbum     string              `json:"first_album"`
	Locations      []string            `json:"locations"`
	ConcertDates   []string            `json:"concert_dates"`
	DatesLocations map[string][]string `json:"dates_locations"`
	IsFavorited    bool
	Followers      int
	Tracks         int
}
