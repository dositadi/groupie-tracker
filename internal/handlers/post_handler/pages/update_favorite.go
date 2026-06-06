package pages

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/pages"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/google/uuid"
)

const (
	sourceUH = "Update favorite handler under pages pkg"
)

func (p *Pages) UpdateFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	favStatus := r.FormValue(utils.FAV_KEY)
	val := r.FormValue(utils.ARTIST_ID_KEY)
	artistId := p.atoi(val)
	page := pages.New(p.logger, w, p.embedded, p.client, r, p.favoriteModel, p.preferenceModel)
	userId := p.getUserId(r)

	switch favStatus {
	case string(pages.FAVORITED):
		exists, err := p.favoriteModel.Exists(artistId)
		if err != nil {
			e := helper.WrapError("Update favorite error", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUH,
			})
			return
		}

		id := uuid.NewString()
		favorite := data.Favorite{
			Id:       id,
			UserId:   userId,
			ArtistId: artistId,
			Status:   false,
		}

		switch exists {
		case false:
			if err := p.favoriteModel.Insert(favorite); err != nil {
				e := helper.WrapError("Update favorite (insert) error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		case true:
			fmt.Println("Entered true Fav")
			status := false
			favUpdate := data.FavoriteUpdate{
				UserId:   userId,
				ArtistId: artistId,
				Status:   &status,
			}
			if err := p.favoriteModel.Update(favUpdate); err != nil {
				e := helper.WrapError("Update favorite error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		}

	case string(pages.NOT_FAVORITED):
		exists, err := p.favoriteModel.Exists(artistId)
		if err != nil {
			e := helper.WrapError("Update favorite error", err)
			p.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUH,
			})
			return
		}

		id := uuid.NewString()
		favorite := data.Favorite{
			Id:       id,
			UserId:   userId,
			ArtistId: artistId,
			Status:   true,
		}

		switch exists {
		case false:
			if err := p.favoriteModel.Insert(favorite); err != nil {
				e := helper.WrapError("Update favorite (insert) error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		case true:
			var status bool = true
			favUpdate := data.FavoriteUpdate{
				UserId:   userId,
				ArtistId: artistId,
				Status:   &status,
			}
			if err := p.favoriteModel.Update(favUpdate); err != nil {
				e := helper.WrapError("Update favorite error", err)
				p.logger.PrintError(e.Error(), map[string]string{
					"Source": sourceUH,
				})
				return
			}
		}
	}
	if err := page.RenderArtistsGrid(pages.FILTER_BY_ID, pages.ASCENDING_ORDER); err != nil {
		e := helper.WrapError("Render favorite button error", err)
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUH,
		})
		return
	}
}

func (p *Pages) atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		p.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": sourceUH,
		})
		panic("Not a valid number")
	}
	return out
}
