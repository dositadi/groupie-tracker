package preferencemodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceUpdate = "Update f(n) under preferencemodel pkg"
)

func (p *PreferenceModel) Update(preference data.PreferenceUpdate) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{})
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
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUpdate,
		})
		return e
	}

	defer tx.Rollback(ctx)

	data, err := p.Get(preference.UserId)
	if err != nil {
		return err
	}

	var filter, sort string
	if preference.Filter != nil {
		filter = *preference.Filter
	}
	if preference.Sort != nil {
		sort = *preference.Sort
	}

	/*
			filter filters NOT NULL DEFAULT 'ID',
		    sort sorts NOT NULL DEFAULT 'ASC',
		    version integer NOT NULL DEFAULT 1,
	*/

	query := "UPDATE preferences SET filter = $1, sort = $2 version = version + 1, updated_at = now() WHERE userId = $3 AND version = $4"

	_, err = tx.Exec(ctx, query, filter, sort, preference.UserId, data.Version)
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
		p.logger.PrintError(e.Error(), map[string]string{
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
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceUpdate,
		})
		return e
	}

	return nil
}
