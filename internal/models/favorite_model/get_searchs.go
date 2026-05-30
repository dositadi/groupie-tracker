package favoritemodel

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
	timeout      = 5
	sourceGetAll = "Get all f(n) under favoritemodel pkg"
)

func (f *FavoriteModel) GetAll(userId string) ([]data.Favorite, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	query := `SELECT artistId FROM favorites WHERE userId = $1`

	rows, err := f.db.Query(ctx, query, userId)
	if err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Query execution error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Query execution error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Query execution error: transaction closed", err)
		case errors.Is(err, pgx.ErrNoRows):
			e = helper.WrapError("No favorite available for this user", err)
		default:
			e = helper.WrapError("Query execution error", err)
		}
		f.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceGetAll,
		})
		return nil, e
	}

	var favorites []data.Favorite

	for rows.Next() {
		var favorite data.Favorite

		if err := rows.Scan(&favorite.ArtistId); err != nil {
			var e error
			switch {
			case errors.Is(err, context.Canceled):
				e = helper.WrapError("Query execution error: context canceled", err)
			case errors.Is(err, context.DeadlineExceeded):
				e = helper.WrapError("Query execution error: deadline exceeded", err)
			case errors.Is(err, pgx.ErrTxClosed):
				e = helper.WrapError("Query execution error: transaction closed", err)
			case errors.Is(err, pgx.ErrNoRows):
				e = helper.WrapError("No favorite available for this user", err)
			default:
				e = helper.WrapError("Query execution error", err)
			}
			f.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceGetAll,
			})
			return nil, e
		}

		favorites = append(favorites, favorite)
	}

	if err := rows.Err(); err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Query execution error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Query execution error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Query execution error: transaction closed", err)
		case errors.Is(err, pgx.ErrNoRows):
			e = helper.WrapError("No favorite available for this user", err)
		default:
			e = helper.WrapError("Query execution error", err)
		}
		f.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceGetAll,
		})
		return nil, e
	}
	return favorites, nil
}
