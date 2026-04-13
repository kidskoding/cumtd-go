// Package cumtd provides a Go client for the Champaign-Urbana Mass Transit
// District (MTD) API v3.
//
// # Getting started
//
// Create a client with your API key and call any method:
//
//	client := cumtd.New("YOUR_API_KEY")
//
//	deps, err := client.GetDepartures(ctx, "STOP1", nil)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, d := range deps {
//	    fmt.Printf("%s — %s\n", d.UniqueID, d.DepartsIn)
//	}
//
// # Authentication
//
// All requests require an API key sent via the X-Api-Key header. Obtain a key
// at https://developer.mtd.org. Note: MTD API v3 uses a header, not a query
// parameter as in v2.
//
// # Error handling
//
// Methods return typed errors. Use errors.As to distinguish them:
//
//	var apiErr *cumtd.APIError
//	var rlErr  *cumtd.RateLimitError
//	var valErr *cumtd.ValidationError
//
//	switch {
//	case errors.As(err, &rlErr):
//	    time.Sleep(retryAfterDuration(rlErr.RetryAfter))
//	case errors.As(err, &apiErr):
//	    log.Printf("API error %d: %s", apiErr.StatusCode, apiErr.Body)
//	case errors.As(err, &valErr):
//	    log.Printf("bad param %s: %s", valErr.Field, valErr.Message)
//	}
//
// # API reference
//
// Full MTD API v3 documentation is available at https://mtd.dev.
package cumtd

import (
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the MTD API v3 base URL.
	DefaultBaseURL = "https://api.mtd.dev"

	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 10 * time.Second

	// Version is the SDK version, sent in the User-Agent header.
	Version = "0.1.0"
)

// Client is the MTD API v3 client. Create one with [New] and reuse it across
// requests. All methods are safe for concurrent use.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	userAgent  string
}

// Option is a functional option for configuring a [Client].
type Option func(*Client)

// New creates a new Client authenticated with the given API key.
// Pass [Option] values to override defaults.
//
//	client := cumtd.New("YOUR_API_KEY",
//	    cumtd.WithTimeout(30*time.Second),
//	)
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

// WithHTTPClient replaces the underlying HTTP client. Use this to set custom
// transport, proxy, or TLS configuration.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// WithBaseURL overrides the API base URL. Useful for testing against a mock
// server or a staging environment.
func WithBaseURL(u string) Option {
	return func(c *Client) { c.baseURL = u }
}

// WithUserAgent appends a custom string to the User-Agent header sent with
// every request.
func WithUserAgent(ua string) Option {
	return func(c *Client) { c.userAgent = ua }
}

// WithTimeout sets the timeout for every HTTP request made by the client.
// Overrides [DefaultTimeout].
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient = &http.Client{Timeout: d}
	}
}
