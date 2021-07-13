package rest

import "time"

type (

	// RequestForCandles holds the necessary information to represent the response for candlestick data.
	RequestForCandles struct {
		Market     string
		Resolution int
		Limit      int
		Start      time.Time
		End        time.Time
	}

	// Candle holds the necessary information to represent a candlestick.
	Candle struct {
		Close     float64   `json:"close"`
		High      float64   `json:"high"`
		Low       float64   `json:"low"`
		Open      float64   `json:"open"`
		StartTime time.Time `json:"startTime"`
		Volume    float64   `json:"volume"`
	}

	// ResponseForCandles holds the necessary information to represent the response for candlestick data
	ResponseForCandles struct {
		BaseResponse
		Result []Candle `json:"result"`
	}

	// RequestForTrades represents a request for trades data.
	RequestForTrades struct {
		Market string
	}

	// Trade holds the necessary information to represent an executed trade
	Trade struct {
		ID          int       `json:"id"`
		Liquidation bool      `json:"liquidation"`
		Price       float64   `json:"price"`
		Side        string    `json:"side"`
		Size        float64   `json:"size"`
		Time        time.Time `json:"time"`
	}

	// ResponseForTrades holds the necessary information to represent the response for trade data
	ResponseForTrades struct {
		BaseResponse
		Result []Trade `json:"result"`
	}

	// RequestForOrderBook represents a request for all markets
	RequestForOrderBook struct {
		Market string
	}

	// OrderBook holds the necessary information to represent an order book
	OrderBook struct {
		Asks [][]float64 `json:"asks"`
		Bids [][]float64 `json:"bids"`
	}

	// ResponseForOrderBook holds the necessary information to represent the response for order book data
	ResponseForOrderBook struct {
		BaseResponse
		Result OrderBook `json:"result"`
	}

	// RequestForMarket represents a request for all markets
	RequestForMarket struct {
		Market string
	}

	// RequestForMarkets represents a request for all markets
	RequestForMarkets struct {
	}

	// Market holds the necessary information to represent a market
	Market struct {
		Name           string      `json:"name"`
		BaseCurrency   interface{} `json:"baseCurrency"`
		QuoteCurrency  interface{} `json:"quoteCurrency"`
		Type           string      `json:"type"`
		Underlying     string      `json:"underlying"`
		Enabled        bool        `json:"enabled"`
		Ask            float64     `json:"ask"`
		Bid            int         `json:"bid"`
		Last           float64     `json:"last"`
		PostOnly       bool        `json:"postOnly"`
		PriceIncrement float64     `json:"priceIncrement"`
		SizeIncrement  float64     `json:"sizeIncrement"`
		Restricted     bool        `json:"restricted"`
	}

	// ResponseForMarkets holds the necessary information to represent the response for markets data
	ResponseForMarkets struct {
		BaseResponse
		Result []Market `json:"result"`
	}

	// ResponseForMarket holds the necessary information to represent the response for a market's data
	ResponseForMarket struct {
		BaseResponse
		Result Market `json:"result"`
	}
)
