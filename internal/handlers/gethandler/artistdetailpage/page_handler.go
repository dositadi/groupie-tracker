package artistdetailpage

import (
	"net/http"

	artistdetail "github.com/dositadi/groupie-tracker/internal/services/pages/artistdetailspage"
)

const (
	sourceD = "Detail page handler f(n) under artistdetailpage"
)

func (d *DetailPage) DetailPageHandler(w http.ResponseWriter, r *http.Request) {
	page := artistdetail.New(d.logger, w, d.embedded, d.client, r)

	if err := page.RenderArtistDetailPage(); err != nil {
		d.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceD,
		})
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
