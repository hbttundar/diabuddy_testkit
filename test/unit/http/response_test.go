package http_test

import (
	"encoding/json"
	testhttp "github.com/hbttundar/diabuddy_testkit/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMustStatus(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	testhttp.MustStatus(t, resp, 200)
}

func TestMustHeader(t *testing.T) {
	resp := &http.Response{Header: http.Header{"X-Test": []string{"value"}}}
	testhttp.MustHeader(t, resp, "X-Test", "value")
}

func TestParseJSONMap(t *testing.T) {
	rr := httptest.NewRecorder()
	json.NewEncoder(rr).Encode(map[string]any{"foo": "bar"})
	resp := rr.Result()

	parsed := testhttp.ParseJSONMap(t, resp.Body)
	assert.Equal(t, "bar", parsed["foo"])
}

func TestMustJSONField(t *testing.T) {
	data := map[string]any{"foo": "bar"}
	testhttp.MustJSONField(t, data, "foo", "bar")
}

func TestAssertPaginationHeaders(t *testing.T) {
	rr := httptest.NewRecorder()
	rr.Header().Set("X-Total", "10")
	rr.Header().Set("X-Page", "1")
	rr.Header().Set("X-Limit", "5")
	rr.Header().Set("X-Total-Pages", "2")

	testhttp.AssertPaginationHeaders(t, rr.Result())
}

func TestAssertSortedBy(t *testing.T) {
	list := []map[string]any{
		{"name": "Alice"},
		{"name": "Bob"},
		{"name": "Charlie"},
	}
	testhttp.AssertSortedBy(t, list, "name")
}

func TestAssertSortedByNumeric(t *testing.T) {
	list := []map[string]any{
		{"score": 1.1},
		{"score": 2.3},
		{"score": 3.9},
	}
	testhttp.AssertSortedByNumeric(t, list, "score")
}
