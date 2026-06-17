package herokuapp


type ArtistInfo struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

type artistMetaData struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type concertDates struct {
	Id           int      `json:"id"`
	ConcertDates []string `json:"dates"`
}

type relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}
