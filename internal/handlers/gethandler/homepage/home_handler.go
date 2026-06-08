package homepage

import (
	"fmt"
	"net/http"

	pages "github.com/dositadi/groupie-tracker/internal/services/pages/home_page"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	TOO_LARGE = "Large request body."
	sourceHH  = "Home Handler function under apppages pkg"
)

func (a *HomePage) HomeHandler(w http.ResponseWriter, r *http.Request) {
	page := pages.New(a.logger, w, a.embedded, a.client, r, a.favoriteModel, a.preferencemodel)

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	if err := r.ParseForm(); err != nil {
		http.Error(w, TOO_LARGE, http.StatusBadRequest)
		a.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceHH,
		})
	}

	pageNo := r.FormValue(utils.PAGE_KEY)
	fmt.Println(pageNo)

	if pageNo == "" {
		if err := page.RenderHomePage(false); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceHH,
			})
		}
	} else {
		if err := page.RenderHomePage(true); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			a.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceHH,
			})
		}
	}
}
