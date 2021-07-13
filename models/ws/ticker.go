package ws

import "time"

type (
	// Ticker represents information about a certain market.
	Ticker struct {
		Bid        float64   `json:"bid"`
		Ask        float64   `json:"ask"`
		LastPrice  float64   `json:"lastPrice"`
		MarkPrice  float64   `json:"markPrice"`
		IndexPrice float64   `json:"indexPrice"`
		Timestamp  time.Time `json:"time"`
	}
)
