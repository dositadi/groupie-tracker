package pages

import (
	"fmt"
	"html/template"
	"strings"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceRS = "Render search f(n) under pages pkg"
)

var currentPageS = 1
var currentIndexS = 0
var countS = 0
var currentSearch = ""

func (p *Pages) RenderSearch() error {
	fs := []string{
		"internal/web/static/partials/pages/home_page_partials.html",
	}

	search := p.request.FormValue(utils.SEARCH_KEY)

	if currentSearch != search {
		currentPageS = 1
		currentIndexS = 0
		countS = 0
		currentSearch = search
	}

	temp, err := template.New("home_page_partials.html").Funcs(p.homePageFunc()).ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRS,
		})
		return e
	}

	var artists []artistapi.ArtistInfo
	var paginatedArtists []artistapi.ArtistInfo
	search = strings.ToLower(search)

	for _, artist := range p.client.GetByIdKey() {
		if strings.Contains(strings.ToLower(artist.FirstAlbum), search) || strings.Contains(strings.ToLower(artist.Name), search) {
			artists = append(artists, artist)
		}
	}

	userFavorites, err := p.getUserFavorites()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceRS,
		})
		return err
	}

	userPreference, err := p.getUserPreference()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceRS,
		})
		return err
	}

	artists = sortArtist(artists, Sort(userPreference.Sort), Filter(userPreference.Filter))

	page := p.request.FormValue(utils.PAGE_KEY)

	if page != "" {
		p := p.atoi(page)
		if p < 0 {
			currentPageS += p
		} else {
			currentPageS = p
		}
	}

	var disableNextButton bool
	var disablePrevButton bool
	artistsLen := len(artists) - 1

	limit := 10
	//totalPages := len(artists) / limit
	offset := (currentPageS - 1) * limit
	fmt.Println(offset)
	currentIndexS = offset + limit

	switch len(artists) {
	case 1:
		paginatedArtists = artists[:]
		countS = 1
		disableNextButton = true
	default:
		if currentIndexS < artistsLen {
			paginatedArtists = artists[offset : offset+limit]
			countS = artistsLen - (artistsLen - (offset + limit))
		} else {
			if offset < artistsLen {
				paginatedArtists = artists[offset:]
				countS = artistsLen + 1
				disableNextButton = true
			} else if offset == artistsLen {
				disableNextButton = true
			}
		}
	}

	if currentPageS == 1 {
		disablePrevButton = true
	}

	data := struct {
		NextPage, PreviousPage, Count, Total                   int
		UserFavorites                                          map[int]data.Favorite
		Artists                                                []artistapi.ArtistInfo
		CurrentFilter, CurrentSort                             string
		FilterSortRoute                                        string
		FilterByName, FilterByCreationDate, FilterByFirstAlbum string
		FilterKey, ArtistIDKey, SearchKey, FavKey, PageKey     string
		SortKey, SortASC, SortDESC                             string
		Favorited, NotFavorited                                string
		FavoriteArtistUrl, Url                                 string
		DisableNextbutton, DisablePrevButton, IsSearch         bool
	}{
		UserFavorites:        userFavorites,
		Artists:              paginatedArtists,
		CurrentFilter:        userPreference.Filter,
		CurrentSort:          userPreference.Sort,
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
		SearchKey:            utils.SEARCH_KEY,
		DisableNextbutton:    disableNextButton,
		DisablePrevButton:    disablePrevButton,
		Count:                countS,
		NextPage:             currentPageS + 1,
		PreviousPage:         currentPageS - 1,
		Total:                len(artists),
		Url:                  utils.ARTIST_SEARCH.String(),
		PageKey:              utils.PAGE_KEY,
		IsSearch:             true,
	}

	if err = temp.ExecuteTemplate(p.responseWriter, "artist-card-main", data); err != nil {
		e := helper.WrapError("Error executing template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRS,
		})
		return e
	}
	return nil
}
