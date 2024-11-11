package yfinance

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
	headers map[string]string
}

type ClientOpt func(*Client)

// NewClient constructs a Client with default fields but allows customization
// through functional options
func NewClient(opts ...ClientOpt) *Client {
	client := &Client{ // default Client implementation
		baseURL: `https://query1.finance.yahoo.com`,
		client:  &http.Client{Timeout: 30 * time.Second},
		headers: make(map[string]string),
	}

	for _, opt := range opts { //override defaults
		opt(client)
	}

	return client
}

// WithBaseURL allows you to customize baseURL for the client
func WithBaseURL(baseURL string) ClientOpt {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithClient allows you to customize the http client
func WithClient(client *http.Client) ClientOpt {
	return func(c *Client) {
		c.client = client
	}
}

// WithHeader allows you to customize multiple headers at once. Last value wins.
func WithHeaders(headers map[string]string) ClientOpt {
	return func(c *Client) {
		for key, value := range headers {
			c.headers[key] = value
		}
	}
}

// Get makes an API call to the yahoo v8 API returning ticker data
func (c *Client) Get(ticker string) (*Chart, error) {
	endpoint := c.baseURL + `/v8/finance/chart/` + ticker
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
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
