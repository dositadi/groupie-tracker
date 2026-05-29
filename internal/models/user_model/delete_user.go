package usermodel

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/jackc/pgx/v5"
)

const (
	sourceDelete = "Delete f(n) under usermodel"
)

var USER_NOT_FOUND = fmt.Errorf("User does not exist.")

func (u *UserModel) Delete(id string) error {
	start := time.Now()

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

	query := `DELETE FROM users WHERE id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), timeOut*time.Second)
	defer cancel()

	_, err = u.db.Exec(ctx, query, id)
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
			"Source": sourceDelete,
		})
		return e
	}

	u.logger.PrintInfo("Deleted a user successfully", map[string]string{
		"Source":     sourceGet2,
		"ID":         id,
		"Total time": time.Since(start).String(),
	})
	return nil
}
