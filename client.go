package yfinance

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type client struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

type ClientOpt func(*client)

// NewClient constructs a client with default fields but allows customization
// through functional options
func NewClient(opts ...ClientOpt) *client {
	client := &client{ // default Client implementation
		BaseURL:    `https://query1.finance.yahoo.com`,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
		Headers:    make(map[string]string),
	}

	for _, opt := range opts { //override defaults
		opt(client)
	}

	return client
}

// WithBaseURL allows you to customize BaseURL for the client
func WithBaseURL(baseURL string) ClientOpt {
	return func(c *client) {
		c.BaseURL = baseURL
	}
}

// WithClient allows you to customize the http client
func WithHTTPClient(httpclient *http.Client) ClientOpt {
	return func(c *client) {
		c.HTTPClient = httpclient
	}
}

// WithHeaders allows you to customize multiple headers at once. Last value wins.
func WithHeaders(headers map[string]string) ClientOpt {
	return func(c *client) {
		for key, value := range headers {
			c.Headers[key] = value
		}
	}
}

// Get makes an API call to the yahoo v8 API returning ticker data
func (c *client) Get(ticker string) (*Chart, error) {
	endpoint := c.BaseURL + `/v8/finance/chart/` + ticker
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Request to %s failed: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Got %d from %s", resp.StatusCode, endpoint)
	}

	var chart Chart
	if err = json.NewDecoder(resp.Body).Decode(&chart); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response: %w", err)
	}

	return &chart, nil
}
