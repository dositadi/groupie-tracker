package favoritemodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceUpdate = "Update f(n) under favoritemodel pkg"
)

func (f *FavoriteModel) Update(fav data.FavoriteUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	tx, err := f.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		var e error
		switch {
		case errors.Is(err, context.Canceled):
			e = helper.WrapError("Transaction begin error: context canceled", err)
		case errors.Is(err, context.DeadlineExceeded):
			e = helper.WrapError("Transaction begin error: deadline exceeded", err)
		case errors.Is(err, pgx.ErrTxClosed):
			e = helper.WrapError("Transaction begin error: transaction closed", err)
		default:
			e = helper.WrapError("Transaction begin error", err)
		}
		f.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUpdate,
		})
		return e
	}

	defer tx.Rollback(ctx)

	data, err := f.Get(fav.ArtistId, fav.UserId)
	if err != nil {
		return err
	}

	if fav.Status != nil {
		data.Status = *fav.Status
	}

	query := "UPDATE favorites SET status = $1, version = version + 1, updated_at = now() WHERE userId = $2 AND artistId = $3 AND version = $4"

	_, err = tx.Exec(ctx, query, data.Status, fav.UserId, fav.ArtistId, data.Version)
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
			"Source": sourceUpdate,
		})
		return e
	}

	if err = tx.Commit(ctx); err != nil {
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
			"Source": sourceUpdate,
		})
		return e
	}

	return nil
}
