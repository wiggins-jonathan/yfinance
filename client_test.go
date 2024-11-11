package yfinance

import (
	"fmt"
	"testing"
)

func TestYahooAPIRealCall(t *testing.T) {
	client := NewClient()

	// Call the Get method to fetch data for the ticker
	ticker := "AAPL"
	chart, err := client.Get(ticker)
	if err != nil {
		t.Fatalf("Error calling api: %v", err)
	}

	if err := chart.Chart.Error; err != nil {
		t.Fatalf("Error received from api: %+v\n", err)
	}

	if len(chart.Chart.Result) < 1 {
		t.Fatalf("Expected non-empty Result data for ticker %s", ticker)
	}

	// Validate the meta data from the response
	result := chart.Chart.Result[0]
	if result.Meta.Symbol != ticker {
		t.Errorf("Expected symbol '%s', got '%s'", ticker, result.Meta.Symbol)
	}
	if result.Meta.RegularMarketPrice <= 0 {
		t.Errorf("Expected positive regularMarketPrice, got %.2f", result.Meta.RegularMarketPrice)
	}
	if result.Meta.LongName == "" {
		t.Errorf("Expected LongName to be non-empty, got '%s'", result.Meta.LongName)
	}
	if len(result.Timestamp) == 0 {
		t.Errorf("Expected non-empty Timestamp")
	}

	// informational
	fmt.Printf("Ticker: %s\n", result.Meta.Symbol)
	fmt.Printf("Market Price: %.2f\n", result.Meta.RegularMarketPrice)
	fmt.Printf("Trading Period: %+v\n", result.Meta.CurrentTradingPeriod)
}
