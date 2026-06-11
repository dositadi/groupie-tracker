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

type Location struct {
	Name string
	Lat  float64
	Lng  float64
}

func (a *ArtistDetail) RenderArtistDetailPage() error {
	fs := []string{
		"internal/web/static/pages/artist_profile.html",
	}

	locations := []Location{
		{Name: "Lagos Island", Lat: 6.4541, Lng: 3.3947},
		{Name: "Ikeja", Lat: 6.6018, Lng: 3.3515},
	}

	jsObject := helper.Marshal(locations)

	id := a.atoi(chi.URLParam(a.request, "id"))

	artistInfo := a.client.GetByIdKey()[id]

	data := struct {
		HomeUrl, ArtistDetailUrl, AllEventsPageUrl string
		ArtistInfo                                 artistapi.ArtistInfo
		AllArtists                                 map[int]artistapi.ArtistInfo
		JsObject                                   template.JS
	}{
		HomeUrl:          utils.HOME.String(),
		ArtistInfo:       artistInfo,
		AllArtists:       a.client.GetByIdKey(),
		ArtistDetailUrl:  utils.ARTIST_DETAILS.String(),
		AllEventsPageUrl: utils.ALL_EVENTS_ROUTES.String(),
		JsObject:         template.JS(jsObject),
	}

	/* New("artist_profile.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...) */

	temp := template.Must(template.New("artist_profile.html").Funcs(a.detailPageFuncMap()).ParseFS(a.embedded.Get(), fs...))

	if err := temp.Execute(a.responseWriter, data); err != nil {
		e := helper.WrapError("Template execute error", err)
		a.logger.PrintError(e.Error(), map[string]string{
			"Status": sourceR,
		})
	}
	return nil
}
