package usermodel

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
	sourceInsert = "Insert user f(n) under usermodel pkg"
)

var USER_EXISTS error = fmt.Errorf("User exists already")

func (u *UserModel) Insert(user data.User) error {
	start := time.Now()
	query := `INSERT INTO users (id, username, email, hashed_password)
	VALUES ($1, $2, $3, $4)`

	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		e := helper.WrapError("Begin transaction error", err)
		u.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return e
	}

	exists, err := u.EmailExists(user.Email)
	if err != nil {
		return err
	}

	if exists {
		e := helper.WrapError("Conflict error", USER_EXISTS)
		u.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return e
	}

	_, err = tx.Exec(ctx, query, user.Id, user.Username, user.Email, user.HashedPassword)
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
		u.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return e
	}

	defer tx.Rollback(ctx)

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
		u.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceInsert,
		})
		return e
	}

	u.logger.PrintInfo("Inserted a user successfully", map[string]string{
		"Source":     sourceInsert,
		"Username":   user.Username,
		"Total time": time.Since(start).String(),
	})
	return nil
}
