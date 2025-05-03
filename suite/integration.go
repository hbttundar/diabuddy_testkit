package suite

import (
	"database/sql"
	db2 "github.com/hbttundar/diabuddy_testkit/db"
	_ "os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hbttundar/diabuddy-api-infra/database"
)

// IntegrationSuite provides a reusable base for DB-integrated test suites.
type IntegrationSuite struct {
	*BaseSuite
	DB database.Connection
	Tx *sql.Tx
}

func NewIntegrationSuite(t *testing.T) *IntegrationSuite {
	s := &IntegrationSuite{
		BaseSuite: NewBaseSuite(t),
	}

	conn, cfg, err := db2.InitializeTestDB(s.Ctx)
	require.NoError(t, err)
	s.DB = conn

	require.NoError(t, db2.RunMigrations(cfg))

	tx, err := conn.DB().Begin()
	require.NoError(t, err)
	s.Tx = tx

	return s
}

func (s *IntegrationSuite) Cleanup() {
	_ = s.Tx.Rollback()
}
