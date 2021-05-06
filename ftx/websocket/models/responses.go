package models

import "time"

type (
	// BaseMessage holds the common attributes of the websocket responses.
	BaseMessage struct {
		Type    string `json:"type"`
		Channel string `json:"channel"`
		Market  string `json:"market"`
	}

	// FutureData holds future data from the markets message
	FutureData struct {
		Name                  string    `json:"name"`
		Underlying            string    `json:"underlying"`
		Type                  string    `json:"type"`
		Expiry                time.Time `json:"expiry"`
		ExpiryDescription     bool      `json:"expiryDescription"`
		Expired               bool      `json:"expired"`
		Perpetual             bool      `json:"perpetual"`
		PostOnly              bool      `json:"postOnly"`
		ImfFactor             float64   `json:"imfFactor"`
		UnderlyingDescription bool      `json:"underlyingDescription"`
		Group                 string    `json:"group"`
		PositionLimitWeight   float64   `json:"positionLimitWeight"`
	}

	// MarketData holds the data from the markets message
	MarketData struct {
		Name           string  `json:"name"`
		Enabled        bool    `json:"enabled"`
		PriceIncrement float64 `json:"priceIncrement"`
		SizeIncrement  float64 `json:"sizeIncrement"`
		Type           string  `json:"type"`
		BaseCurrency   string  `json:"baseCurrency"`
		QuoteCurrency  string  `json:"quoteCurrency"`
		Underlying     string  `json:"underlying"`
		Restricted     bool    `json:"restricted"`
		FutureData     `json:"future"`
	}

	// MarketMessage
	MarketMessage struct {
		BaseMessage
		Data MarketData `json:"data"`
	}

	// TradeData holds the data from the markets message
	TradeData struct {
		Side        string  `json:"side"`
		Size        float64 `json:"size"`
		Price       float64 `json:"price"`
		Liquidation bool    `json:"liquidation"`
	}

	// TradeMessage
	TradeMessage struct {
		BaseMessage
		Data []TradeData `json:"data"`
	}

	// OrderBookData holds the data from the markets message
	OrderBookData struct {
	}

	// OrderBookMessage
	OrderBookMessage struct {
		BaseMessage
		Data OrderBookData `json:"data"`
	}

	// TickerData holds the data from the markets message
	TickerData struct {
		Bid       float64 `json:"bid"`
		Ask       float64 `json:"ask"`
		Last      float64 `json:"last"`
		Timestamp float64 `json:"timestamp"`
	}

	// TickerMessage
	TickerMessage struct {
		BaseMessage
		Data TickerData `json:"data"`
	}
)
