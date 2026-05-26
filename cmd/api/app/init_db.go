package app

import (
	"context"
	"fmt"
	"os"
	"time"

	groupietracker "github.com/dositadi/groupie-tracker"
	"github.com/dositadi/groupie-tracker/internal/helper"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
)

func (a *App) initDB() {
	db, err := a.connectToDB()
	if err != nil {
		a.logger.PrintFatal(err.Error(), map[string]string{
			"Source": "Initialize db f(n)",
		})
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()

	if err = db.Ping(ctx); err != nil {
		a.logger.PrintFatal(err.Error(), map[string]string{
			"Source": "Init db f(n)",
		})
		os.Exit(1)
	}

	err = a.migrate()
	if err != nil {
		os.Exit(1)
	}

	a.logger.PrintInfo("Database Initialized successfully", map[string]string{
		"Source": "Init db f(n)",
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
				"Source": "Connect to db f(n)",
			})
			return nil, e
		}
		time.Sleep(5 * time.Second)
	}

	a.logger.PrintInfo("Database connected successfully", map[string]string{
		"Source": "Connect to db f(n)",
	})
	return db, nil
}

func (a *App) migrate() error {
	embedded := groupietracker.NewMigrations()

	migratefiles, err := iofs.New(embedded.Get(), "migrations")
	if err != nil {
		e := helper.WrapError("iofs driver failed", err)
		a.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Migrate f(n)",
		})
		return e
	}

	m, err := migrate.NewWithSourceInstance("iofs", migratefiles, a.config.dsn)
	if err != nil {
		e := helper.WrapError("Migrate new instance failed", err)
		a.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Migrate f(n)",
		})
		return e
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		e := helper.WrapError("Migration failed", err)
		a.logger.PrintFatal(e.Error(), map[string]string{
			"Source": "Migrate f(n)",
		})
		return e
	}

	a.logger.PrintInfo("Migration is successful", map[string]string{
		"Source": "Migrate f(n)",
	})
	return nil
}
