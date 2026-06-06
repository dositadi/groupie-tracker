package pages

import (
	"fmt"
	"html/template"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/data"
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

	fmt.Println(filterBy," : ",sortBy)

	temp, err := template.New("home_page_partials.html").Funcs(p.homePageFunc()).ParseFS(p.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Error creating template", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceAG,
		})
		return e
	}

	userFavorites, err := p.getUserFavorites()
	if err != nil {
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceAG,
		})
		return err
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
		UserFavorites                                          map[int]data.Favorite
		Artists                                                []artistapi.ArtistInfo
		CurrentFilter, CurrentSort                             string
		FilterSortRoute                                        string
		FilterByName, FilterByCreationDate, FilterByFirstAlbum string
		FilterKey,
		ArtistIDKey string
		SortKey, SortASC, SortDESC                         string
		FavoriteArtistUrl, FavKey, Favorited, NotFavorited string
	}{
		UserFavorites:        userFavorites,
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
