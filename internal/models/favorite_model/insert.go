package favoritemodel

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceInsert = "Insert f(n) under favoritemodel"
)

var CONFLICT_ERR = fmt.Errorf("Favorite exists already")

func (f *FavoriteModel) Insert(favorite data.Favorite) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	tx, err := f.db.BeginTx(ctx, pgx.TxOptions{})
	defer tx.Rollback(ctx)
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
			"Source": sourceGetAll,
		})
		return e
	}

	exists, err := f.Exists(favorite.ArtistId)
	if err != nil {
		return err
	}
	if exists {
		f.logger.PrintError(CONFLICT_ERR.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return CONFLICT_ERR
	}

	query := "INSERT INTO favorites (userId, artistId) VALUES ($1, $2)"

	_, err = tx.Exec(ctx, query, favorite.UserId, favorite.ArtistId)
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
			"Source": sourceGetAll,
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
			"Source": sourceGetAll,
		})
		return e
	}

	return nil
}
