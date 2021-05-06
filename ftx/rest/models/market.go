package models

import "time"

// Candle holds the necessary information to represent a candlestick
type Candle struct {
	Close     float64   `json:"close"`
	High      float64   `json:"high"`
	Low       float64   `json:"low"`
	Open      float64   `json:"open"`
	StartTime time.Time `json:"startTime"`
	Volume    float64   `json:"volume"`
}

// Trade holds the necessary information to represent an executed trade
type Trade struct {
	ID          int       `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

// OrderBook holds the necessary information to represent an order book
type OrderBook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// Market holds the necessary information to represent a market
type Market struct {
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

// Future holds the necessary information to represent a future
type Future struct {
	Ask                 int       `json:"ask"`
	Bid                 float64   `json:"bid"`
	Change1H            int       `json:"change1h"`
	Change24H           int       `json:"change24h"`
	ChangeBod           int       `json:"changeBod,omitempty"`
	VolumeUsd24H        int       `json:"volumeUsd24h,omitempty"`
	Volume              float64   `json:"volume,omitempty"`
	Description         string    `json:"description"`
	Enabled             bool      `json:"enabled"`
	Expired             bool      `json:"expired"`
	Expiry              time.Time `json:"expiry"`
	Index               float64   `json:"index"`
	ImfFactor           float64   `json:"imfFactor"`
	Last                int       `json:"last"`
	LowerBound          float64   `json:"lowerBound"`
	Mark                float64   `json:"mark"`
	Name                string    `json:"name"`
	Perpetual           bool      `json:"perpetual"`
	PositionLimitWeight float64   `json:"positionLimitWeight,omitempty"`
	PostOnly            bool      `json:"postOnly"`
	PriceIncrement      float64   `json:"priceIncrement"`
	SizeIncrement       float64   `json:"sizeIncrement"`
	Underlying          string    `json:"underlying"`
	UpperBound          float64   `json:"upperBound"`
	Type                string    `json:"type"`
}

// FutureStats holds the necessary information to represent the future's statistics
type FutureStats struct {
	Volume                   float64   `json:"volume"`
	NextFundingRate          float64   `json:"nextFundingRate"`
	NextFundingTime          time.Time `json:"nextFundingTime"`
	ExpirationPrice          float64   `json:"expirationPrice"`
	PredictedExpirationPrice float64   `json:"predictedExpirationPrice"`
	StrikePrice              float64   `json:"strikePrice"`
	OpenInterest             float64   `json:"openInterest"`
}

// FundingRate holds the necessary information to represent the funding rate
type FundingRate struct {
	Future string    `json:"future"`
	Rate   float64   `json:"rate"`
	Time   time.Time `json:"time"`
}
