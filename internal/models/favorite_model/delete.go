package favoritemodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceDeleteAll = "Delete all f(n) under favoritemodel pkg"
	sourceDelete    = "Delete f(n) under favoritemodel pkg"
)

func (f *FavoriteModel) DeleteAll(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	query := "DELETE FROM favorites WHERE userId = $1"

	_, err := f.db.Exec(ctx, query, userId)
	if err != nil {
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
			"Source": sourceDeleteAll,
		})
		return e
	}

	return nil
}

func (f *FavoriteModel) Delete(userId string, artistId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()
	query := "DELETE FROM favorites WHERE userId = $1 AND artistId = $2"

	_, err := f.db.Exec(ctx, query, userId)
	if err != nil {
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
			"Source": sourceDeleteAll,
		})
		return e
	}

	return nil
}
