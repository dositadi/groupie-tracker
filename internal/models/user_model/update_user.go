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

/*
id uuid NOT NULL PRIMARY KEY,
    username text NOT NULL,
    email citext NOT NULL UNIQUE,
    hashed_password bytea NOT NULL,
    version integer NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
*/

const (
	sourceUpdate = "Update f(n) under usermodel"
)

var EMAIL_EXISTS = fmt.Errorf("This email exists already")

func (u *UserModel) Update(id string, info data.UpdateUser) error {
	start := time.Now()
	user, err := u.GetWithID(id)
	if err != nil {
		u.logger.PrintError(err.Error(), map[string]string{
			"Source": sourceUpdate,
		})
		return err
	}

	if info.HashedPassword != nil {
		user.HashedPassword = info.HashedPassword
	}

	if info.Username != nil {
		user.Username = *info.Username
	}

	if info.Email != nil {
		exists, err := u.EmailExists(*info.Email)
		if err != nil {
			u.logger.PrintError(err.Error(), map[string]string{
				"Source": sourceUpdate,
			})
			return err
		}

		if exists {
			e := helper.WrapError("Conflict", EMAIL_EXISTS)
			u.logger.PrintError(e.Error(), map[string]string{
				"Source": sourceUpdate,
			})
			return e
		}

		user.Email = *info.Email
	}

	exists, err := u.IDExists(id)
	if err != nil {
		return err
	}

	if !exists {
		u.logger.PrintError(USER_NOT_FOUND.Error(), map[string]string{
			"Source": sourceDelete,
		})
		return USER_NOT_FOUND
	}

	query := `UPDATE users SET username = $1, email = $2, hashed_password = $3, version = version + 1, updated_at = now() WHERE id = $4 AND version = $5`
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	_, err = u.db.Exec(ctx, query, user.Username, user.Email, user.HashedPassword, user.Id, user.Version)
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
			"Source": sourceUpdate,
		})
		return e
	}

	u.logger.PrintInfo("Updated a user detail successfully", map[string]string{
		"Source":     sourceUpdate,
		"Username":   user.Username,
		"Total time": time.Since(start).String(),
	})
	return nil
}
