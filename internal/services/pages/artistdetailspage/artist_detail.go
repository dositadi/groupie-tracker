package artistdetail

import (
	"html/template"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/utils"
	"github.com/go-chi/chi/v5"
)

const (
	sourceR = "Render artist detail page f(n) under artistdetail pkg"
)

func (a *ArtistDetail) RenderArtistDetailPage() error {
	fs := []string{
		"internal/web/static/pages/artist_profile.html",
	}

	id := a.atoi(chi.URLParam(a.request, "id"))

	artistInfo := a.client.GetByIdKey()[id]

	temp, err := template.New("artist_profile.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template create error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceR,
		})
	}

	data := struct {
		HomeUrl,ArtistDetailUrl    string
		ArtistInfo artistapi.ArtistInfo
		AllArtists map[int]artistapi.ArtistInfo
	}{
		HomeUrl:    utils.HOME.String(),
		ArtistInfo: artistInfo,
		AllArtists: a.client.GetByIdKey(),
		ArtistDetailUrl: utils.ARTIST_DETAILS.String(),
	}

	if err = temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceR,
		})
	}
	return nil
}
