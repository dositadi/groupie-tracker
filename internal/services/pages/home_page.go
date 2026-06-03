package pages

import (
	"html/template"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

type Filter string
type Sort string
type Favorite string

const (
	sourceRHome = "Render Home page f(n) under pages pkg"
	// Filters
	FILTER_BY_ID            Filter = "ID"
	FILTER_BY_NAME          Filter = "NAME"
	FILTER_BY_CREATION_DATE Filter = "CREATION DATE"
	FILTER_BY_FIRST_ALBUM   Filter = "FIRST ALBUM"

	// Sort orders
	ASCENDING_ORDER  Sort = "ASC"
	DESCENDING_ORDER Sort = "DESC"

	// Favorite
	FAVORITED     Favorite = "true"
	NOT_FAVORITED Favorite = "false"
)

func (p *Pages) RenderHomePage(filterBy Filter, sortBy Sort) error {
	fs := []string{
		"internal/web/static/pages/home_page.html",
		"internal/web/static/partials/pages/home_page_partials.html",
	}

	temp, err := template.New("home_page.html").ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRHome,
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
		SortKey, SortASC, SortDESC string
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
		NotFavorited:         string(FAVORITED),
	}

	if err := temp.Execute(p.responseWriter, data); err != nil {
		e := helper.WrapError("Error executing template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceRHome,
		})
		return e
	}
	return nil
}
