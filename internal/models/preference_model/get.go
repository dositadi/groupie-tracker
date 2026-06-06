package preferencemodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

/*
id uuid NOT NULL PRIMARY KEY,
    userId uuid NOT NULL,
    artistId uuid NOT NULL,
    version integer NOT NULL DEFAULT 1,
*/

const (
	//sourceGetAll = "Get all f(n) under preferencemodel pkg"
	sourceGet = "Get f(n) under preferencemodel pkg"
)

func (p *PreferenceModel) Get(userId string) (data.Preference, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := "SELECT filter, sort, version FROM preferences WHERE userId = $1"

	row := p.db.QueryRow(ctx, query, userId)

	var pref data.Preference

	if err := row.Scan(&pref.Filter, &pref.Sort, &pref.Version); err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Query execution error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Query execution error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Query execution error: transaction closed", err)
		default:
			e = helper.WrapError("Query execution error", err)
		}
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceGet,
		})
		return data.Preference{}, e
	}
	return pref, nil
}
