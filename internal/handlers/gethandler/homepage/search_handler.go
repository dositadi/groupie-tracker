package homepage

import (
	"net/http"

	pages "github.com/dositadi/groupie-tracker/internal/services/pages/home_page"
)

const (
	sourceSH = "Sear"
)

func (p *HomePage) SearchHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.New(p.logger, w, p.embedded, p.client, r, p.favoriteModel, p.preferencemodel)

	if err := page.RenderSearch(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceSH,
		})
	}
}
