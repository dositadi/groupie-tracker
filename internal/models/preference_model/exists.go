package preferencemodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceExists = "Exists f(n) under preferencemodel"
)

func (p *PreferenceModel) Exists(userId string) (bool, error) {
	//  Remember to create an index for the artist id
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM preferences WHERE userId = $1)`

	row := p.db.QueryRow(ctx, query, userId)

	var exists bool

	if err := row.Scan(&exists); err != nil {
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
			"Source": sourceExists,
		})
		return false, e
	}
	return exists, nil
}
