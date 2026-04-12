package testutil

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// MockRoute defines a single mock API route.
type MockRoute struct {
	StatusCode  int
	Body        []byte
	AssertQuery func(t *testing.T, q url.Values)
}

// NewMockServer creates an httptest.Server that serves the given routes.
// Unregistered paths respond with 404.
func NewMockServer(t *testing.T, routes map[string]MockRoute) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route, ok := routes[r.URL.Path]
		if !ok {
			http.Error(w, `{"error":{"message":"not found"}}`, http.StatusNotFound)
			return
		}
		if route.AssertQuery != nil {
			route.AssertQuery(t, r.URL.Query())
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(route.StatusCode)
		_, _ = w.Write(route.Body)
	}))
	t.Cleanup(srv.Close)
	return srv
}

// MustReadFixture reads a fixture file from internal/testutil/fixtures/.
// It locates fixtures relative to this source file, so it works from any
// test package in the module.
func MustReadFixture(t *testing.T, name string) []byte {
	t.Helper()
	_, thisFile, _, _ := runtime.Caller(0)
	path := filepath.Join(filepath.Dir(thisFile), "fixtures", name)
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("MustReadFixture: %v", err)
	}
	return b
}
