package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type RequestOption func(t *testing.T, req *http.Request)

func WithJSONBody(body any) RequestOption {
	return func(t *testing.T, req *http.Request) {
		b, err := json.Marshal(body)
		require.NoError(t, err)
		req.Body = io.NopCloser(bytes.NewBuffer(b))
		req.Header.Set("Content-Type", "application/json")
	}
}

func WithHeader(key, value string) RequestOption {
	return func(t *testing.T, req *http.Request) {
		req.Header.Set(key, value)
	}
}

func WithBearerToken(token string) RequestOption {
	return WithHeader("Authorization", "Bearer "+token)
}

func Send(t *testing.T, method, url string, opts ...RequestOption) *http.Response {
	t.Helper()
	req, err := http.NewRequest(method, url, nil)
	require.NoError(t, err)

	for _, opt := range opts {
		opt(t, req)
	}

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	return resp
}

func Get(t *testing.T, url string, opts ...RequestOption) *http.Response {
	return Send(t, http.MethodGet, url, opts...)
}

func Post(t *testing.T, url string, opts ...RequestOption) *http.Response {
	return Send(t, http.MethodPost, url, opts...)
}

func Patch(t *testing.T, url string, opts ...RequestOption) *http.Response {
	return Send(t, http.MethodPatch, url, opts...)
}

func Delete(t *testing.T, url string, opts ...RequestOption) *http.Response {
	return Send(t, http.MethodDelete, url, opts...)
}
