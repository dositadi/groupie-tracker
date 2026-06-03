package data

/*
 id uuid NOT NULL PRIMARY KEY,
    userId uuid NOT NULL,
    artistId uuid NOT NULL,
    version integer NOT NULL DEFAULT 1,
*/

type Favorite struct {
	Id       string
	UserId   string
	ArtistId int
	Status   bool
	Version  int
}

type FavoriteUpdate struct {
	Id, UserId string
	ArtistId   int
	Status     *bool
}
