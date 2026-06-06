package apppages

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/pages"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceFS = "Filter-Sort handler f(n) under apppages pkg"
)

func (p *Pages) FilterSortHandler(w http.ResponseWriter, r *http.Request) {
	filter := r.FormValue(utils.FILTER_KEY)
	sort := r.FormValue(utils.SORT_KEY)

	page := pages.New(p.logger, w, p.embedded, p.client, r, p.favoriteModel)

	if err := page.RenderArtistsGrid(pages.Filter(filter), pages.Sort(sort)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		p.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceFS,
		})
	}
}
