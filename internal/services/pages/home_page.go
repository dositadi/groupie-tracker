package pages

import (
	"html/template"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceRHome = "Render Home page f(n) under pages pkg"
)

var currentPage = 1
var currentIndex = 0
var count = 0

func (p *Pages) RenderHomePage(partial bool) error {
	fs := []string{
		"internal/web/static/pages/home_page.html",
		"internal/web/static/partials/pages/home_page_partials.html",
	}

	userFavorites, err := p.getUserFavorites()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return err
	}

	userPreference, err := p.getUserPreference()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return err
	}

	temp, err := template.New("home_page.html").Funcs(p.homePageFunc()).ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return e
	}

	var artists []artistapi.ArtistInfo
	var paginatedArtists []artistapi.ArtistInfo

	artists = sortArtist(mapToSlice(p.client.GetByIdKey()), Sort(userPreference.Sort), Filter(userPreference.Filter))

	page := p.request.FormValue(utils.PAGE_KEY)

	if page != "" {
		p := p.atoi(page)
		if p < 0 {
			currentPage += p
		} else {
			currentPage = p
		}
	}

	var disableNextButton bool
	var disablePrevButton bool
	artistsLen := len(artists) - 1

	limit := 10
	//totalPages := len(artists) / limit
	offset := (currentPage - 1) * limit
	currentIndex = offset + limit

	if currentIndex < artistsLen {
		paginatedArtists = artists[offset : offset+limit]
		count = artistsLen - (artistsLen - (offset + limit))
	} else {
		if offset < artistsLen {
			paginatedArtists = artists[offset:]
			count = artistsLen + 1
			disableNextButton = true
		} else if offset == artistsLen {
			disableNextButton = true
		}
	}

	if currentPage == 1 {
		disablePrevButton = true
	}

	data := struct {
		NextPage, PreviousPage, Count, Total                   int
		UserFavorites                                          map[int]data.Favorite
		Artists                                                []artistapi.ArtistInfo
		CurrentFilter, CurrentSort                             string
		FilterSortRoute                                        string
		FilterByName, FilterByCreationDate, FilterByFirstAlbum string
		FilterKey, ArtistIDKey, SearchKey, PageKey             string
		SortKey, SortASC, SortDESC                             string
		FavoriteArtistUrl, FavKey, Favorited, NotFavorited     string
		SearchUrl, Url                                         string
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
		SearchUrl:            utils.ARTIST_SEARCH.String(),
		SearchKey:            utils.SEARCH_KEY,
		Url:                  utils.HOME.String(),
		PageKey:              utils.PAGE_KEY,
		NextPage:             currentPage + 1,
		PreviousPage:         currentPage - 1,
		Count:                count,
		Total:                len(artists),
		DisableNextbutton:    disableNextButton,
		DisablePrevButton:    disablePrevButton,
		IsSearch:             false,
	}

	if partial {
		if err = temp.ExecuteTemplate(p.responseWriter, "artist-card-main", data); err != nil {
			e := helper.WrapError("Error executing template", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceRHome,
			})
			return e
		}
	} else {
		if err = temp.Execute(p.responseWriter, data); err != nil {
			e := helper.WrapError("Error executing template", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceRHome,
			})
			return e
		}
	}

	return nil
}

/* func (p *Pages) isHTMXRequest() bool {
	return p.request.Header.Get("HX-Request") == "true"
} */
