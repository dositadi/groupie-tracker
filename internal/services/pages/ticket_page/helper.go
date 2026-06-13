package ticketpage

import (
	"html/template"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"unicode"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/services/ordercache"
	"github.com/dositadi/groupie-tracker/internal/utils"
)

func (t *TicketPage) getTicketPrice(ticketType string) float64 {
	switch ticketType {
	case string(ordercache.GENERAL):
		return math.Round(float64(ordercache.GENERAL_AMT)*100) / 100
	case string(ordercache.RESERVED):
		return math.Round(float64(ordercache.RESERVED_AMT)*100) / 100
	case string(ordercache.VIP):
		return math.Round(float64(ordercache.VIP_AMT)*100) / 100
	default:
		return 0
	}
}

func (t *TicketPage) atoi(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		t.logger.PrintError("Atoi conversion error: Not a valid number", map[string]string{
			"Source": sourceT,
		})
		panic("Not a valid number")
	}
	return out
}

func (t *TicketPage) detailPageFuncMap() template.FuncMap {
	return template.FuncMap{
		/* "RangeEvents": func(relations map[string][]string) {

		}, */
		"GetDay": func(date string) string {
			out := strings.Split(date, "-")
			return out[0]
		},
		"GetMonth": func(date string) string {
			out := strings.Split(date, "-")
			num := t.atoi(out[1])

			switch num {
			case 1:
				return "JAN"
			case 2:
				return "FEB"
			case 3:
				return "MAR"
			case 4:
				return "APR"
			case 5:
				return "MAY"
			case 6:
				return "JUN"
			case 7:
				return "JUL"
			case 8:
				return "AUG"
			case 9:
				return "SEPT"
			case 10:
				return "OCT"
			case 11:
				return "NOV"
			case 12:
				return "DEC"
			default:
				return "JAN"
			}
		},
		"GetYear": func(date string) string {
			out := strings.Split(date, "-")
			return out[2]
		},
		"CleanCityName": func(city string) string {
			city = strings.ReplaceAll(city, "-", " ")
			city = strings.ReplaceAll(city, "_", "-")
			citySlice := strings.Split(city, " ")

			for i, c := range citySlice {
				v := []rune(c)
				citySlice[i] = string(unicode.ToUpper(v[0])) + string(v[1:])
			}

			return strings.Join(citySlice, ", ")
		},
		"RandomValues": func() int {
			return rand.Intn(500)
		},
		"TotalTicketAmount": func(ticketPrice float64, qty int) float64 {
			amt := ticketPrice * float64(qty)
			return math.Round(amt*100) / 100
		},
		"TotalBookingFee": func(fee float64, qty int) float64 {
			amt := fee * float64(qty)
			return math.Round(amt*100) / 100
		},
		"VatAmount": func(vatRate int, totalTicketAmount, totalBookingFee float64) float64 {
			totalPrice := totalBookingFee + totalTicketAmount
			vatAmount := totalPrice * (float64(vatRate) / (100 + float64(vatRate)))
			return math.Round(vatAmount*100) / 100
		},
		"GrandTotal": func(totalTicketAmount, totalBookingFee, vatAmount float64) float64 {
			total := totalBookingFee + totalTicketAmount + vatAmount
			return math.Round(total*100) / 100
		},
	}
}

func (t *TicketPage) getUser() data.User {
	val := t.request.Context().Value(utils.USER_ID_KEY)

	if user, ok := val.(data.User); ok {
		return user
	}
	return data.User{}
}
