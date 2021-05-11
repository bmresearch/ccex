package websocket

const (
	/*
		General
	*/
	Url = "wss://stream.bybit.com/realtime"

	/*
		Message types
	*/
	Subscribed   = "subscribed"
	Unsubscribed = "unsubscribed"
	Snapshot     = "snapshot"
	Delta        = "delta"

	/*
		Public websocket channels
	*/
	OrderBook    = "orderBook"
	OrderBook25  = "orderBookL2_25"
	OrderBook200 = "orderBook_200"
	Instrument   = "instrument_info"
	Trades       = "trade"

	/*
		Private websocket channels
	*/

	/*
		Side
	*/
	Buy  = "Buy"
	Sell = "Sell"
)
