package websocket

const (
	/*
		General
	*/
	Url = "wss://ftx.com/ws/"

	/*
		Message types
	 */
	Subscribed = "subscribed"
	Unsubscribed = "unsubscribed"
	Snapshot = "partial"
	Delta = "update"

	/*
		Public websocket channels
	*/
	Ticker    = "ticker"
	Trades    = "trades"
	Orderbook = "orderbook"

	/*
		Private websocket channels
	*/
	Fills  = "fills"
	Orders = "orders"
)
