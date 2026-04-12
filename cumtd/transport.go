package cumtd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type apiErrBody struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// get performs a GET request, injects auth headers, and decodes dst from the
// response envelope's data field.
func (c *Client) get(ctx context.Context, path string, params url.Values, dst any) error {
	u := c.baseURL + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return &RateLimitError{RetryAfter: resp.Header.Get("Retry-After")}
	}
	if resp.StatusCode != http.StatusOK {
		return &APIError{StatusCode: resp.StatusCode, Body: string(body)}
	}

	var raw struct {
		Data  json.RawMessage `json:"data"`
		Error *apiErrBody     `json:"error,omitempty"`
	}
	if err := json.Unmarshal(body, &raw); err != nil {
		return fmt.Errorf("cumtd: decode envelope: %w", err)
	}
	if raw.Error != nil {
		return &APIError{StatusCode: resp.StatusCode, Body: raw.Error.Message}
	}
	return json.Unmarshal(raw.Data, dst)
}
