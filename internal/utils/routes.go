package utils

type Route string

func (r Route) String() string {
	return string(r)
}

const (
	LOGIN             Route = Route("/auth/session")
	REGISTER          Route = Route("/auth/registration")
	HOME              Route = Route("/artists")
	FAVORITES         Route = Route("/artists/favorites")
	ARTIST_DETAILS    Route = Route("/artists/detail")
	ARTIST_SEARCH     Route = Route("/artists/search") //?query=
	EVENTS            Route = Route("/artists/concert")
	FAVORITE          Route = Route("/artists/favorite")
	FILTER_SORT_ROUTE Route = Route("/artists/filter-sort")
	ALL_EVENTS_ROUTES Route = Route("/artists/all-events")
)
