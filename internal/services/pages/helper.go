package pages

import (
	"html/template"
	"maps"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"unicode"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
)

// This function converts a map to a slice
func mapToSlice[T comparable, K artistapi.ArtistInfo](artists map[T]K) []K {
	return slices.Collect(maps.Values(artists))
}

func sortArtist(artists []artistapi.ArtistInfo, sort Sort, filter Filter) []artistapi.ArtistInfo {
	switch filter {
	case FILTER_BY_ID:
		slices.SortStableFunc(artists, func(a, b artistapi.ArtistInfo) int {
			if a.Id < b.Id {
				return -1
			} else if a.Id > b.Id {
				return 1
			} else {
				return 0
			}
		})

		switch sort {
		case ASCENDING_ORDER:
			return artists
		case DESCENDING_ORDER:
			slices.Reverse(artists)
			return artists
		}

	case FILTER_BY_NAME:
		slices.SortStableFunc(artists, func(a, b artistapi.ArtistInfo) int {
			if a.Name < b.Name {
				return -1
			} else if a.Name > b.Name {
				return 1
			} else {
				return 0
			}
		})

		switch sort {
		case ASCENDING_ORDER:
			return artists
		case DESCENDING_ORDER:
			slices.Reverse(artists)
			return artists
		}

	case FILTER_BY_CREATION_DATE:
		slices.SortStableFunc(artists, func(a, b artistapi.ArtistInfo) int {
			if a.CreationDate < b.CreationDate {
				return -1
			} else if a.CreationDate > b.CreationDate {
				return 1
			} else {
				return 0
			}
		})

		switch sort {
		case ASCENDING_ORDER:
			return artists
		case DESCENDING_ORDER:
			slices.Reverse(artists)
			return artists
		}

	case FILTER_BY_FIRST_ALBUM:
		slices.SortStableFunc(artists, func(a, b artistapi.ArtistInfo) int {
			if a.FirstAlbum < b.FirstAlbum {
				return -1
			} else if a.FirstAlbum > b.FirstAlbum {
				return 1
			} else {
				return 0
			}
		})

		switch sort {
		case ASCENDING_ORDER:
			return artists
		case DESCENDING_ORDER:
			slices.Reverse(artists)
			return artists
		}
	}
	return artists
}

func (p *Pages) getUserFavorites() (map[int]data.Favorite, error) {
	favorites, err := p.favoriteModel.GetAll(p.getUserId())
	if err != nil {
		e := helper.WrapError("Favorites fetch error", err)
		return nil, e
	}

	favMap := make(map[int]data.Favorite)

	for _, favorite := range favorites {
		favMap[favorite.ArtistId] = favorite
	}
	return favMap, nil
}

func (p *Pages) getUserPreference() (data.Preference, error) {
	pref, err := p.preferenceModel.Get(p.getUserId())
	if err != nil {
		e := helper.WrapError("Preference fetch error", err)
		return data.Preference{}, e
	}
	return pref, nil
}

func (p *Pages) homePageFunc() template.FuncMap {
	return template.FuncMap{
		"GetLocation": func(s []string) string {
			return s[0]
		},
		"CleanText": func(s string) string {
			s = strings.ReplaceAll(s, "_", " ")
			s = strings.ReplaceAll(s, "-", " ")
			s = strings.ToLower(s)
			sl := strings.Fields(s)

			for i, w := range sl {
				rn := []rune(w)
				rn[0] = unicode.ToUpper(rn[0])
				sl[i] = string(rn)
			}

			return strings.Join(sl, " ")
		},
		"CheckFav": func(artist artistapi.ArtistInfo, favorites map[int]data.Favorite) bool {
			if favorites == nil {
				return artist.IsFavorited
			}
			if status, ok := favorites[artist.Id]; ok {
				return status.Status
			}
			return artist.IsFavorited
		},
		"RandomValues": func() int {
			return rand.Intn(500)
		},
	}
}

func (p *Pages) atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		p.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": sourceRHome,
		})
		panic("Not a valid number")
	}
	return out
}
