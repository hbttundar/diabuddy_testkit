package db

import (
	"database/sql"
	baseerrors "errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/hbttundar/diabuddy-api-config/config/dbconfig"
	"github.com/hbttundar/diabuddy-api-config/util/resolver/rootpath"
	"log"
	"path/filepath"
)

// RunMigrations applies all up migrations.
func RunMigrations(config *dbconfig.DBConfig) error {
	base, apiErr := rootpath.NewRootPathResolver().Resolve("./")
	if apiErr != nil {
		return fmt.Errorf("failed to resolve root path: %w", apiErr)
	}

	path := fmt.Sprintf("file://%s/migrations", base)
	connStr, apiErr := config.ConnectionString()
	if apiErr != nil {
		return fmt.Errorf("failed to build DB connection string: %w", apiErr)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open DB connection: %w", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to initialize migration: %w", err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && !baseerrors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}

// RollbackMigrations rolls back all migrations (use with care).
func RollbackMigrations(config *dbconfig.DBConfig) error {
	base, apiErr := rootpath.NewRootPathResolver().Resolve("./")
	if apiErr != nil {
		log.Printf("Failed to resolve root path: %v", apiErr)
	}
	path := fmt.Sprintf("file://%s/migrations", filepath.Join(base))
	connStr, apiErr := config.ConnectionString()
	if apiErr != nil {
		return fmt.Errorf("error building connection string: %w", apiErr)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error in connection to postgres database")
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating postgres driver: %w", err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}
	defer m.Close()

	if err := m.Down(); err != nil {
		return fmt.Errorf("migration down failed: %w", err)
	}
	return nil
}
