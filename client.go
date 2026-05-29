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
		Headers:    map[string]string{"User-Agent": randomUserAgent()},
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

// randomUserAgent returns a random user agent string to identify our client
func randomUserAgent() string {
	agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
	}

	var seedOnce sync.Once
	seedOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	return agents[rand.Intn(len(agents))]
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
