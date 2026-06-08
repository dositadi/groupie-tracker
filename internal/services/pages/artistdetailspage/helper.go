package artistdetail

import (
	"html/template"
	"strconv"
	"strings"
	"unicode"
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
			citySlice := strings.Split(city, " ")

			for i, c := range citySlice {
				v := []rune(c)
				citySlice[i] = string(unicode.ToUpper(v[0])) + string(v[1:])
			}

			return strings.Join(citySlice, " ")
		},
		"GetLocations": func (locations []string) string {
			return locations[0]
		},
		
	}
}
