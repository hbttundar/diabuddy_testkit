package http_test

import (
	testhttp "github.com/hbttundar/diabuddy_testkit/http"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	_ "strings"
	"testing"
)

func TestGet(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		w.WriteHeader(http.StatusOK)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	resp := testhttp.Get(t, ts.URL)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestPostWithJSONBody(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.Contains(t, string(bodyBytes), "email")
		assert.NoError(t, err)
		w.WriteHeader(http.StatusCreated)
	})

	ts := httptest.NewServer(handler)
	defer ts.Close()

	resp := testhttp.Post(t, ts.URL, testhttp.WithJSONBody(map[string]any{
		"email": "test@example.com",
	}))
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
