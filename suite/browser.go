package suite

import (
	"net/http/httptest"
	"testing"

	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"github.com/stretchr/testify/require"
	"net/http"
)

// BrowserSuite extends IntegrationSuite with a test HTTP router and client.
type BrowserSuite struct {
	*IntegrationSuite
	Router *router.Router
	Server *httptest.Server
	Client *http.Client
}

func NewBrowserSuite(t *testing.T, buildRouter func(bs *BrowserSuite) *router.Router) *BrowserSuite {
	s := NewIntegrationSuite(t)
	bs := &BrowserSuite{IntegrationSuite: s}

	r := buildRouter(bs)
	require.NotNil(t, r, "buildRouter must return a non-nil *router.Router")

	handler, ok := r.Adapter().(http.Handler)
	require.True(t, ok, "router.Adapter() must implement http.Handler")
	server := httptest.NewServer(handler)
	bs.Router = r
	bs.Server = server
	bs.Client = server.Client()

	return bs
}

func (s *BrowserSuite) Cleanup() {
	s.IntegrationSuite.Cleanup()
}

func WithBrowserTestSuite(t *testing.T, buildRouter func(bs *BrowserSuite) *router.Router, fn func(suite *BrowserSuite)) {
	s := NewBrowserSuite(t, buildRouter)
	defer s.Cleanup()
	fn(s)
}
