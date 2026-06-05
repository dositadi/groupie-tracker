package pages

import (
	"html/template"
	"strings"
	"unicode"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceAG = "Render artist grid f(n) under pages pkg"
)

func (p *Pages) RenderArtistsGrid(filterBy Filter, sortBy Sort) error {
	fs := []string{
		"internal/web/static/partials/pages/home_page_partials.html",
	}

	funcMap := template.FuncMap{
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
	}

	temp, err := template.New("home_page_partials.html").Funcs(funcMap).ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceAG,
		})
		return e
	}

	var artists []artistapi.ArtistInfo

	switch filterBy {
	case FILTER_BY_ID:
		artists = sortArtists(p.client.GetByIdKey(), sortBy)
	case FILTER_BY_NAME:
		artists = sortArtists(p.client.GetByName(), sortBy)
	case FILTER_BY_FIRST_ALBUM:
		artists = sortArtists(p.client.GetByFirstAlbum(), sortBy)
	case FILTER_BY_CREATION_DATE:
		artists = sortArtists(p.client.GetByCreationDate(), sortBy)
	default:
		filterBy = FILTER_BY_ID
		sortBy = ASCENDING_ORDER
		artists = sortArtists(p.client.GetByIdKey(), sortBy)
	}

	data := struct {
		Artists                                                []artistapi.ArtistInfo
		CurrentFilter, CurrentSort                             string
		FilterSortRoute                                        string
		FilterByName, FilterByCreationDate, FilterByFirstAlbum string
		FilterKey,
		ArtistIDKey string
		SortKey, SortASC, SortDESC                         string
		FavoriteArtistUrl, FavKey, Favorited, NotFavorited string
	}{
		Artists:              artists,
		CurrentFilter:        string(filterBy),
		CurrentSort:          string(sortBy),
		FilterSortRoute:      utils.FILTER_SORT_ROUTE.String(),
		FilterByName:         string(FILTER_BY_NAME),
		FilterByCreationDate: string(FILTER_BY_CREATION_DATE),
		FilterByFirstAlbum:   string(FILTER_BY_FIRST_ALBUM),
		FilterKey:            utils.FILTER_KEY,
		SortKey:              utils.SORT_KEY,
		SortASC:              string(ASCENDING_ORDER),
		SortDESC:             string(DESCENDING_ORDER),
		FavoriteArtistUrl:    utils.FAVORITE.String(),
		FavKey:               utils.FAV_KEY,
		Favorited:            string(FAVORITED),
		NotFavorited:         string(NOT_FAVORITED),
		ArtistIDKey:          utils.ARTIST_ID_KEY,
	}

	if err = temp.ExecuteTemplate(p.responseWriter, "artist-card-main", data); err != nil {
		e := helper.WrapError("Error executing template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceAG,
		})
		return e
	}

	return nil
}
