package db

import (
	"context"
	"github.com/hbttundar/diabuddy-api-config/config/dbconfig"
	"github.com/hbttundar/diabuddy-api-config/config/envmanager"
	"github.com/hbttundar/diabuddy-api-infra/database"
)

func InitializeTestDB(ctx context.Context) (database.Connection, *dbconfig.DBConfig, error) {
	envMgr, err := envmanager.NewEnvManager(envmanager.WithEnvironment("test"), envmanager.WithUseDefault(true))
	if err != nil {
		return nil, nil, err
	}

	cfg, err := dbconfig.NewDBConfig(envMgr, dbconfig.WithType(dbconfig.Postgres), dbconfig.WithDsnParameters(map[string]string{"sslmode": "disable"}))
	if err != nil {
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
