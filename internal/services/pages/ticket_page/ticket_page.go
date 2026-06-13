package ticketpage

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	artistapi "github.com/dositadi/groupie-tracker/internal/client/artist_api"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

const (
	sourceT = "Render ticket page f(n) under ticketpage pkg"
)

func (t *TicketPage) RenderTicketPage() error {
	fs := []string{
		"internal/web/static/pages/ticket_purchase_page.html",
	}

	path := t.request.FormValue(utils.PATH_KEY)
	artistId := t.atoi(t.request.FormValue(utils.ARTIST_ID_KEY))
	fmt.Println(artistId)
	date := t.request.FormValue(utils.DATE_KEY)
	location := t.request.FormValue(utils.LOCATION_KEY)
	user := t.getUser()

	// Add the user's order to the cache
	ordercache.Set(user.Id, location, artistId, string(ordercache.GENERAL))

	var artistInfo artistapi.ArtistInfo

	if val, ok := t.client.GetByIdKey()[artistId]; ok {
		artistInfo = val
	} else {
		err := errors.New("Artist ID does not exist")
		t.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceT,
		})
		http.Error(t.responseWriter, err.Error(), http.StatusBadRequest)
	}

	booking, exist := ordercache.Get(user.Id, location, artistId)
	if !exist {
		t.logger.PrintError(NOT_FOUND.Error(), map[string]string{
			"Source": sourcePa,
		})
		return NOT_FOUND
	}

	fmt.Println("passed: ", booking)

	data := struct {
		TicketType                                             string
		Quantity                                               int
		TicketPrice                                            float64
		BookingFee                                             float64
		VatValue                                               int
		ArtistInfo                                             artistapi.ArtistInfo
		Location, Date                                         string
		PreviousPageUrl, TicketTypeUrl                         string
		ArtistId                                               int
		ArtistIdKey, DateKey, LocationKey, TicketTypeKey       string
		GeneralTicket, VipTicket, ReserveTicket                string
		GeneralTicketPrice, VipTicketPrice, ReserveTicketPrice float64
	}{
		TicketType:         booking.TicketType,
		Quantity:           booking.Quantity,
		TicketPrice:        t.getTicketPrice(booking.TicketType),
		BookingFee:         float64(ordercache.BOOKING_FEE),
		VatValue:           int(ordercache.VAT),
		GeneralTicket:      string(ordercache.GENERAL),
		VipTicket:          string(ordercache.VIP),
		ReserveTicket:      string(ordercache.RESERVED),
		GeneralTicketPrice: float64(ordercache.GENERAL_AMT),
		VipTicketPrice:     float64(ordercache.VIP_AMT),
		ReserveTicketPrice: float64(ordercache.RESERVED_AMT),
		TicketTypeKey:      utils.TICKET_TYPE_KEY,
		ArtistIdKey:        utils.ARTIST_ID_KEY,
		DateKey:            utils.DATE_KEY,
		LocationKey:        utils.LOCATION_KEY,
		ArtistId:           artistId,
		TicketTypeUrl:      utils.TicketType.String(),
		ArtistInfo:         artistInfo,
		Location:           location,
		Date:               date,
		PreviousPageUrl:    path,
	}

	temp, err := template.New("ticket_purchase_page.html").Funcs(t.detailPageFuncMap()).ParseFS(t.embedded.Get(), fs...)
	if err != nil {
		e := helper.WrapError("Template parse error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceT,
		})
		return e
	}

	if err := temp.Execute(t.responseWriter, data); err != nil {
		e := helper.WrapError("Template execution error", err)
		t.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceT,
		})
		return e
	}
	return nil
}
