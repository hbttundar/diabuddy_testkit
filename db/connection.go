package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/hbttundar/diabuddy-api-config/config/dbconfig"
	"github.com/hbttundar/diabuddy-api-config/config/envmanager"
	"github.com/hbttundar/diabuddy-api-infra/database"
	"log"
	"strings"
)

func EnsureTestDatabase(config *dbconfig.DBConfig) error {
	log.Println("Start creating Test database...")
	dsn, err := config.ConnectionString()
	if err != nil {
		return fmt.Errorf("failed to get connection string: %w", err)
	}

	// switch to the default "postgres" database
	defaultDSN := strings.Replace(dsn, config.Get("DB_DATABASE"), "postgres", 1)
	dbConn, sqlErr := sql.Open("postgres", defaultDSN)
	if sqlErr != nil {
		return fmt.Errorf("failed to connect to default database: %w", sqlErr)
	}
	defer dbConn.Close()

	dbName := config.Get("DB_DATABASE")
	if dbName == "" {
		return errors.New("database name is not specified in the configuration")
	}

	// Check if the database exists
	var exists bool
	row := dbConn.QueryRowContext(context.Background(),
		"SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", dbName)
	if err := row.Scan(&exists); err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}

	if !exists {
		createDbSQL := fmt.Sprintf("CREATE DATABASE %s;", dbName)
		if _, sqlErr = dbConn.ExecContext(context.Background(), createDbSQL); sqlErr != nil &&
			!strings.Contains(sqlErr.Error(), "already exists") {
			return fmt.Errorf("failed to create test database: %w", sqlErr)
		}
	}

	log.Println("Test database is ready.")
	return nil
}

func InitializeTestDB(ctx context.Context) (database.Connection, *dbconfig.DBConfig, error) {
	envMgr, err := envmanager.NewEnvManager(envmanager.WithEnvironment("test"), envmanager.WithUseDefault(true))
	if err != nil {
		return nil, nil, err
	}

	cfg, err := dbconfig.NewDBConfig(envMgr, dbconfig.WithType(dbconfig.Postgres), dbconfig.WithDsnParameters(map[string]string{"sslmode": "disable"}))
	if err != nil {
		return nil, nil, err
	}

	if err := EnsureTestDatabase(cfg); err != nil {
		return nil, nil, err
	}

	conn, err := database.NewPostgresConnection(database.WithPostgresConfig(cfg))
	if err != nil {
		return nil, nil, err
	}
	if cnnErr := conn.Open(ctx); cnnErr != nil {
		return nil, nil, cnnErr
	}

	return conn, cfg, nil
}
