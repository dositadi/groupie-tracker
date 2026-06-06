package preferencemodel

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
	sourceInsert = "Insert f(n) under preferencemodel"
	timeout      = 5
)

/*
CREATE TABLE IF NOT EXISTS preference (
    id uuid NOT NULL PRIMARY KEY,
    userId uuid NOT NULL,
    filter filters NOT NULL DEFAULT 'ID',
    sort sorts NOT NULL DEFAULT 'ASC',
    version integer NOT NULL DEFAULT 1,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updateAt TIMESTAMP WITH TIME ZONE DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);
*/

var CONFLICT_ERR = fmt.Errorf("Preference exists already")

func (p *PreferenceModel) Insert(preference data.Preference) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	tx, err := p.db.BeginTx(ctx, pgx.TxOptions{})
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
		p.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return e
	}

	exists, err := p.Exists(preference.UserId)
	if err != nil {
		return err
	}
	if exists {
		p.logger.PrintError(CONFLICT_ERR.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return CONFLICT_ERR
	}

	query := "INSERT INTO preferences (id, userId, filter, sort) VALUES ($1, $2, $3, $4)"

	_, err = tx.Exec(ctx, query, preference.Id, preference.UserId, preference.Filter, preference.Sort)
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
			"Source": sourceInsert,
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
			"Source": sourceInsert,
		})
		return e
	}

	return nil
}
