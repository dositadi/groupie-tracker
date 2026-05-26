package artistapi

type Artist struct {
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

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"` // Further get request
}

type ConcertDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
