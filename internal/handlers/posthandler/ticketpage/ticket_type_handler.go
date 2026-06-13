package ticketpage

import (
	"net/http"

	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	ticketpage "github.com/dositadi/groupie-tracker/internal/services/pages/ticket_page"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceT = "Ticket type handler f(n) under ticketpage"
)

func (t *TicketPage) TicketTypeHandler(w http.ResponseWriter, r *http.Request) {
	artistId := t.atoi(r.FormValue(utils.ARTIST_ID_KEY), sourceT)
	location := r.FormValue(utils.LOCATION_KEY)
	ticketType := r.FormValue(utils.TICKET_TYPE_KEY)
	user := t.getUserId(r)

	ordercache.Set(user.Id, location, artistId, ticketType)

	partial := ticketpage.New(t.logger, w, t.embedded, t.client, r)

	if err := partial.RenderTicketPagePartials(user.Id, location, artistId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceT,
		})
	}
}
