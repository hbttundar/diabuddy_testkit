package http

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// MustStatus asserts the response status code matches expected.
func MustStatus(t *testing.T, resp *http.Response, expected int) {
	t.Helper()
	require.Equal(t, expected, resp.StatusCode)
}

// MustHeader asserts the response has a specific header value.
func MustHeader(t *testing.T, resp *http.Response, key, expected string) {
	t.Helper()
	actual := resp.Header.Get(key)
	require.Equal(t, expected, actual, "header mismatch for %s", key)
}

// ParseJSONMap parses the response body as a map[string]any.
func ParseJSONMap(t *testing.T, body io.ReadCloser) map[string]any {
	t.Helper()
	defer body.Close()

	var result map[string]any
	err := json.NewDecoder(body).Decode(&result)
	require.NoError(t, err)
	return result
}

// MustJSONField asserts a specific field in a parsed JSON response.
func MustJSONField(t *testing.T, data map[string]any, key string, expected any) {
	t.Helper()
	val, ok := data[key]
	require.True(t, ok, "missing key: %s", key)
	require.Equal(t, expected, val)
}

// AssertPaginationHeaders checks X-Total, X-Page, X-Limit headers.
func AssertPaginationHeaders(t *testing.T, resp *http.Response) {
	t.Helper()
	for _, key := range []string{"X-Total", "X-Page", "X-Limit", "X-Total-Pages"} {
		val := resp.Header.Get(key)
		require.NotEmpty(t, val, "pagination header missing: %s", key)
	}
}

// AssertSortedBy checks that a slice of maps is sorted by a given string key.
func AssertSortedBy(t *testing.T, list []map[string]any, key string) {
	t.Helper()
	sorted := make([]string, len(list))
	for i, item := range list {
		sorted[i] = item[key].(string)
	}
	copyCheck := append([]string(nil), sorted...)
	sort.Strings(copyCheck)
	require.Equal(t, copyCheck, sorted, "response is not sorted by %s", key)
}

// AssertSortedByNumeric checks that a slice of maps is sorted by a given numeric key.
func AssertSortedByNumeric(t *testing.T, list []map[string]any, key string) {
	t.Helper()
	numbers := make([]float64, len(list))
	for i, item := range list {
		num, ok := item[key].(float64)
		require.True(t, ok, "value for key %s is not numeric", key)
		numbers[i] = num
	}
	copyCheck := append([]float64(nil), numbers...)
	sort.Float64s(copyCheck)
	require.Equal(t, copyCheck, numbers, "response is not numerically sorted by %s", key)
}
