// Package cumtd provides a Go client for the CUMTD (MTD) API v3.
// See https://mtd.dev for API documentation.
package cumtd

import (
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the MTD API v3 base URL.
	DefaultBaseURL = "https://api.mtd.dev/api/v3"
	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 10 * time.Second
	// Version is the SDK version.
	Version = "0.1.0"
)

// Client is the MTD API client. All methods are safe for concurrent use.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	userAgent  string
}

// Option configures a Client.
type Option func(*Client)

// New creates a new Client with the given API key and options.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:  apiKey,
		baseURL: DefaultBaseURL,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		userAgent: "cumtd-go/" + Version,
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithHTTPClient replaces the underlying HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithBaseURL overrides the API base URL (useful for testing).
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = u }
}

// WithUserAgent sets a custom User-Agent header.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient = &http.Client{Timeout: d}
	}
}
