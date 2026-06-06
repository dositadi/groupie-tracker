package apppages

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/pages"
)

const (
	sourceSH = "Sear"
)

func (p *Pages) SearchHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.New(p.logger, w, p.embedded, p.client, r, p.favoriteModel, p.preferencemodel)

	if err := page.RenderSearch(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceSH,
		})
	}
}
