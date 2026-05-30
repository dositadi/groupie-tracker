package data

/*
id uuid NOT NULL PRIMARY KEY,
    search text NOT NULL,
    userId uuid NOT NULL,
*/

type Searches struct {
	Id     string
	Search string
	UserId string
}
