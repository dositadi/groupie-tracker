package artistdetail

import (
	"fmt"
	"html/template"
	"maps"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"unicode"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
)

func (a *ArtistDetail) atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		a.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": sourceR,
		})
		panic("Not a valid number")
	}
	return out
}

func (a *ArtistDetail) detailPageFuncMap() template.FuncMap {
	return template.FuncMap{
		/* "RangeEvents": func(relations map[string][]string) {

		}, */
		"GetDay": func(date string) string {
			out := strings.Split(date, "-")
			return out[0]
		},
		"GetMonth": func(date string) string {
			out := strings.Split(date, "-")
			num := a.atoi(out[1])

			switch num {
			case 1:
				return "JAN"
			case 2:
				return "FEB"
			case 3:
				return "MAR"
			case 4:
				return "APR"
			case 5:
				return "MAY"
			case 6:
				return "JUN"
			case 7:
				return "JUL"
			case 8:
				return "AUG"
			case 9:
				return "SEPT"
			case 10:
				return "OCT"
			case 11:
				return "NOV"
			case 12:
				return "DEC"
			default:
				return "JAN"
			}
		},
		"GetYear": func(date string) string {
			out := strings.Split(date, "-")
			return out[2]
		},
		"CleanCityName": func(city string) string {
			city = strings.ReplaceAll(city, "-", " ")
			city = strings.ReplaceAll(city, "_", "-")
			citySlice := strings.Split(city, " ")

			for i, c := range citySlice {
				v := []rune(c)
				citySlice[i] = string(unicode.ToUpper(v[0])) + string(v[1:])
			}

			return strings.Join(citySlice, ", ")
		},
		"GetLocations": func(locations []string) string {
			return locations[0]
		},
		"Artists": func(relations map[string][]string) map[string][]string {
			vals := slices.Collect(maps.Keys(relations))
			slices.Sort(vals)

			limit := 3

			if len(relations) < limit {
				limit = len(relations)
			}

			keys := vals[:3]
			fmt.Println(keys)

			out := make(map[string][]string)

			for _, key := range keys {
				out[key] = relations[key]
			}
			fmt.Println(out)
			return out
		},
		"RandomValues": func() int {
			return rand.Intn(500)
		},
		"SimilarArtists": func(id int, artists map[int]artistapi.ArtistInfo) []artistapi.ArtistInfo {
			keys := slices.Collect(maps.Keys(artists))
			values := slices.Collect(maps.Values(artists))
			slices.Sort(keys)
			slices.SortStableFunc(values, func(a, b artistapi.ArtistInfo) int {
				if a.Id < b.Id {
					return -1
				} else if a.Id > b.Id {
					return 1
				} else {
					return 0
				}
			})

			idx := slices.Index(keys, id)

			var out []artistapi.ArtistInfo

			if (len(values[:idx+1]) + 3) > len(values)-1 {
				out = values[idx-3 : idx]
			} else {
				out = values[idx+1 : idx+4]
			}
			return out
		},
	}
}
