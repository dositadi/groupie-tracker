package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func (a *App) initDB() {
	db, err := a.connectToDB()
	if err != nil {
		a.logger.PrintFatal(err.Error(), map[string]string{
			"Source": "Initialize db function",
		})
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()

	if err = db.Ping(ctx); err != nil {
		a.logger.PrintFatal(err.Error(), map[string]string{
			"Source":  "Initialize db function",
			"Problem": "Database ping failed",
		})
		os.Exit(1)
	}

	a.logger.PrintInfo("Database Initialized successfully", map[string]string{
		"Source": "Initialize db function",
	})

	a.db = db
}

func (a *App) connectToDB() (*pgx.Conn, error) {
	var db *pgx.Conn
	for i := 0; i <= 5; i++ {
		var err error
		db, err = pgx.Connect(context.Background(), a.config.dsn)
		if err != nil {
			e := fmt.Errorf("Database connect err: %w", err)
			a.logger.PrintError(e.Error(), map[string]string{
				"Source": "Connect to db function",
			})
			return nil, e
		}
		time.Sleep(5 * time.Second)
	}
	a.logger.PrintInfo("Database connected successfully", map[string]string{
		"Source": "Connect to db function",
	})
	return db, nil
}
