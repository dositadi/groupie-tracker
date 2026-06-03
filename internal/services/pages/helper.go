package pages

import (
	"maps"
	"slices"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
)

// This function converts a map to a slice
func mapToSlice[T comparable, K artistapi.ArtistInfo](artists map[T]K) []K {
	return slices.Collect(maps.Values(artists))
}

type keyType interface {
	~int | string
}

// This function sorts an array of type ArtistsInfo
func sortArtists[T keyType, K artistapi.ArtistInfo, L Sort](artistsMap map[T]K, order L) []K {
	keys := slices.Collect(maps.Keys(artistsMap))
	slices.Sort(keys)

	var sorted []K

	switch any(order).(Sort) {
	case ASCENDING_ORDER:
		for _, key := range keys {
			artist := artistsMap[key]

			sorted = append(sorted, artist)
		}
	case DESCENDING_ORDER:
		slices.Reverse(keys)

		for _, key := range keys {
			artist := artistsMap[key]

			sorted = append(sorted, artist)
		}
	default:
		for _, key := range keys {
			artist := artistsMap[key]

			sorted = append(sorted, artist)
		}
	}
	return sorted
}
