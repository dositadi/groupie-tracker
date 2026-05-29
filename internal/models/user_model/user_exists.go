package usermodel

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceExists1 = "Email exists f(n) under usermodel pkg"
	sourceExists2 = "ID exists f(n) under usermodel pkg"
)

// This function checks if a particular user exists in the database.
func (u *UserModel) EmailExists(email string) (bool, error) {
	start := time.Now()
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`
	exists := false
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	row := u.db.QueryRow(ctx, query, email)

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
		u.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceExists1,
		})
		return false, e
	}

	u.logger.PrintInfo("Checked if a user exists successfully", map[string]string{
		"Source":     sourceExists1,
		"Email":      email,
		"Status":     strconv.FormatBool(exists),
		"Total time": time.Since(start).String(),
	})
	return exists, nil
}

func (u *UserModel) IDExists(id string) (bool, error) {
	start := time.Now()
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)`
	exists := false
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	row := u.db.QueryRow(ctx, query, id)

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
		u.logger.PrintError(e.Error(), map[string]string{
			"Source": sourceExists2,
		})
		return false, e
	}

	u.logger.PrintInfo("Checked if a user exists successfully", map[string]string{
		"Source":     sourceExists2,
		"ID":         id,
		"Status":     strconv.FormatBool(exists),
		"Total time": time.Since(start).String(),
	})
	return exists, nil
}
