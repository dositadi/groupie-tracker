package usermodel

import (
	"context"
	"errors"
	"time"

	"github.com/dositadi/groupie-tracker/internal/data"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceGet1 = "Get with email f(n) under usermodel"
	sourceGet2 = "Get with id f(n) under usermodel"
)

// This function Gets a particular user from the database.
func (u *UserModel) GetWithEmail(email string) (data.User, error) {
	start := time.Now()

	var user data.User
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	query := `SELECT id, username, hashed_password WHERE email = $1`
	row := u.db.QueryRow(ctx, query, email)

	err := row.Scan(&user.Id, &user.Username, &user.HashedPassword)
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
			"Source": sourceGet1,
		})
		return data.User{}, e
	}

	u.logger.PrintInfo("Fetched a user successfully", map[string]string{
		"Source":     sourceGet1,
		"Username":   user.Username,
		"Total time": time.Since(start).String(),
	})
	return user, nil
}

// This function Gets a particular user from the database.
func (u *UserModel) GetWithID(id string) (data.User, error) {
	start := time.Now()
	var user data.User
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	query := `SELECT username, email, version FROM users WHERE id = $1`
	row := u.db.QueryRow(ctx, query, id)

	if err := row.Scan(&user.Username, &user.Email, &user.Version); err != nil {
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
			"Source": sourceGet2,
		})
		return data.User{}, e
	}

	u.logger.PrintInfo("Fetched a user successfully", map[string]string{
		"Source":     sourceGet2,
		"Username":   user.Username,
		"Total time": time.Since(start).String(),
	})
	return user, nil
}
