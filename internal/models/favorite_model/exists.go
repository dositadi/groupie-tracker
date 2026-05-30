package favoritemodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceExists = "Exists f(n) under favoritesmodel"
)

func (f *FavoriteModel) Exists(artistId int) (bool, error) {
	//  Remember to create an index for the artist id
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := `SELECT EXISTS (SELECT 1 FROM favorites WHERE artistId = $1)`

	row := f.db.QueryRow(ctx, query, artistId)

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
		f.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceGetAll,
		})
		return false, e
	}
	return exists, nil
}
