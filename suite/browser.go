package suite

import (
	"testing"

	infrahttp "github.com/hbttundar/diabuddy-api-infra/http/request"
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"github.com/stretchr/testify/require"
	"net/http"
)

// BrowserSuite extends IntegrationSuite with a test HTTP router and client.
type BrowserSuite struct {
	*IntegrationSuite
	Router *router.Router
	Client *http.Client
}

func NewBrowserSuite(t *testing.T, buildRouter func() *router.Router) *BrowserSuite {
	s := &BrowserSuite{
		IntegrationSuite: NewIntegrationSuite(t),
	}

	s.Router = buildRouter()
	s.Client = infrahttp.DefaultHTTPClient()

	require.NotNil(t, s.Router)
	require.NotNil(t, s.Client)

	return s
}

func (s *BrowserSuite) Cleanup() {
	s.IntegrationSuite.Cleanup()
}

func WithBrowserTestSuite(t *testing.T, buildRouter func() *router.Router, fn func(suite *BrowserSuite)) {
	s := NewBrowserSuite(t, buildRouter)
	defer s.Cleanup()
	fn(s)
}
