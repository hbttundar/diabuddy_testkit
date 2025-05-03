package suite

import (
	"testing"

	"github.com/gin-gonic/gin"
	infrahttp "github.com/hbttundar/diabuddy-api-infra/http"
	"github.com/stretchr/testify/require"
	"net/http"
)

// BrowserSuite extends IntegrationSuite with a test HTTP router and client.
type BrowserSuite struct {
	*IntegrationSuite
	Router *gin.Engine
	Client *http.Client
}

func NewBrowserSuite(t *testing.T, setupRoutes func(r *gin.Engine)) *BrowserSuite {
	s := &BrowserSuite{
		IntegrationSuite: NewIntegrationSuite(t),
	}

	s.Router = infrahttp.SetupRouter(setupRoutes)
	s.Client = infrahttp.DefaultHTTPClient()

	require.NotNil(t, s.Router)
	require.NotNil(t, s.Client)

	return s
}

func (s *BrowserSuite) Cleanup() {
	s.IntegrationSuite.Cleanup()
}
