package cumtd_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/kidskoding/cumtd-go/cumtd"
)

func TestNew_defaults(t *testing.T) {
	c := cumtd.New("key")
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestWithBaseURL(t *testing.T) {
	c := cumtd.New("key", cumtd.WithBaseURL("http://localhost:9999"))
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestWithHTTPClient(t *testing.T) {
	hc := &http.Client{Timeout: 5 * time.Second}
	c := cumtd.New("key", cumtd.WithHTTPClient(hc))
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestWithUserAgent(t *testing.T) {
	c := cumtd.New("key", cumtd.WithUserAgent("my-app/1.0"))
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}

func TestWithTimeout(t *testing.T) {
	c := cumtd.New("key", cumtd.WithTimeout(30*time.Second))
	if c == nil {
		t.Fatal("expected non-nil client")
	}
}
