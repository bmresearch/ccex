package ws

import "time"

type (
	// Trades is defined as an helper type for a slice of trades.
	Trades []Trade

	// Trade represents a trade in the real-time trades feed.
	Trade struct {
		Side                string    `json:"side"`
		Size                float64   `json:"size"`
		Price               float64   `json:"price"`

		/*
			ProvidesLiquidation is used to define if the underlying
			real-time trades feed provides information on whether a
			trade was a liquidation event or not. In cases where
			this is false, the underlying `Liquidation` attribute
			may not necessarily represent if the trade was indeed
			part of a liquidation event.
		 */
		ProvidesLiquidation bool      `json:"providesLiquidation"`
		Liquidation         bool      `json:"liquidation"`
		Time                time.Time `json:"time"`
	}
)
