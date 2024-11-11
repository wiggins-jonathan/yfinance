package yfinance

// Chart represents the entire JSON structure returned from yahoo v8
type Chart struct {
	Chart struct {
		Result []Result `json:"result"`
		Error  any      `json:"error"`
	} `json:"chart"`
}

// Result contains the main data for each Result in the response
type Result struct {
	Meta       Meta       `json:"meta"`
	Timestamp  []int64    `json:"timestamp"`
	Indicators Indicators `json:"indicators"`
}

// Meta holds meta data for the result
type Meta struct {
	Currency             string            `json:"currency"`
	Symbol               string            `json:"symbol"`
	ExchangeName         string            `json:"exchangeName"`
	FullExchangeName     string            `json:"fullExchangeName"`
	InstrumentType       string            `json:"instrumentType"`
	FirstTradeTimestamp  int               `json:"firstTradeDate"`
	RegularMarketTime    int               `json:"regularMarketTime"`
	HasPrePostMarketData bool              `json:"hasPrePostMarketData"`
	GMTOffset            int               `json:"gmtoffset"`
	Timezone             string            `json:"timezone"`
	ExchangeTimezoneName string            `json:"exchangeTimezoneName"`
	RegularMarketPrice   float64           `json:"regularMarketPrice"`
	FiftyTwoWeekHigh     float64           `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64           `json:"fiftyTwoWeekLow"`
	RegularMarketDayHigh float64           `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64           `json:"regularMarketDayLow"`
	RegularMarketVolume  int               `json:"regularMarketVolume"`
	LongName             string            `json:"longName"`
	ShortName            string            `json:"shortName"`
	ChartPreviousClose   float64           `json:"chartPreviousClose"`
	PreviousClose        float64           `json:"previousClose"`
	Scale                int               `json:"scale"`
	PriceHint            int               `json:"priceHint"`
	CurrentTradingPeriod TradingPeriod     `json:"currentTradingPeriod"`
	TradingPeriods       [][]TradingPeriod `json:"tradingPeriods"`
	DataGranularity      string            `json:"dataGranularity"`
	Range                string            `json:"range"`
	ValidRanges          []string          `json:"validRanges"`
}

// TradingPeriod represents a trading period's time data
type TradingPeriod struct {
	Timezone  string `json:"timezone"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
	GMTOffset int    `json:"gmtoffset"`
}

// Indicators contains market data like quotes
type Indicators struct {
	Quote []Quote `json:"quote"`
}

// Quote represents market data for a given period
type Quote struct {
	Open   []float64 `json:"open"`
	Close  []float64 `json:"close"`
	High   []float64 `json:"high"`
	Volume []int     `json:"volume"`
	Low    []float64 `json:"low"`
}
