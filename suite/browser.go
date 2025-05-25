package suite

import (
	"github.com/hbttundar/diabuddy-api-infra/http/router"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

// BrowserSuite extends IntegrationSuite with a test HTTP router and client.
type BrowserSuite struct {
	*IntegrationSuite
	Router *router.Router
	Server *httptest.Server
	Client *http.Client
}

func NewBrowserSuite(t *testing.T, buildRouter func(bs *BrowserSuite) *router.Router) *BrowserSuite {
	is := NewIntegrationSuite(t)

	bs := &BrowserSuite{IntegrationSuite: is}

	testRouter := buildRouter(bs)
	require.NotNil(t, testRouter, "buildRouter must return a non-nil *router.Router")

	bs.Router = testRouter
	bs.Server = httptest.NewServer(testRouter)
	bs.Client = bs.Server.Client()
	require.NotNil(t, bs.Server, "Server should never be nil")
	require.NotNil(t, bs.Client, "Client should never be nil")

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
