package websocket

const (
	/*
		General
	*/
	Url = "wss://ftx.com/ws/"

	/*
		Public websocket channels
	*/
	Ticker    = "ticker"
	Markets   = "markets"
	Trades    = "trades"
	Orderbook = "orderbook"

	/*
		Private websocket channels
	*/
	Fills  = "fills"
	Orders = "orders"
)
