package pages

import (
	"html/template"
	"maps"
	"slices"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceRHome = "Render Home page f(n) under pages pkg"
)

func (p *Pages) RenderHomePage() error {
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

	artists = sortSearchedArtist(slices.Collect(maps.Values(p.client.GetByIdKey())), Sort(userPreference.Sort), Filter(userPreference.Filter))

	data := struct {
		UserFavorites                                          map[int]data.Favorite
		Artists                                                []artistapi.ArtistInfo
		CurrentFilter, CurrentSort                             string
		FilterSortRoute                                        string
		FilterByName, FilterByCreationDate, FilterByFirstAlbum string
		FilterKey, ArtistIDKey, SearchKey                      string
		SortKey, SortASC, SortDESC                             string
		FavoriteArtistUrl, FavKey, Favorited, NotFavorited     string
		SearchUrl                                              string
	}{
		UserFavorites:        userFavorites,
		Artists:              artists,
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
	}

	if err = temp.Execute(p.responseWriter, data); err != nil {
		e := helper.WrapError("Error executing template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return e
	}

	return nil
}

/* func (p *Pages) isHTMXRequest() bool {
	return p.request.Header.Get("HX-Request") == "true"
} */
