package examples

import "github.com/murlokito/ccex/models/ws"

type Data struct {
	Trades map[string][]ws.Trade
	Ticker map[string]ws.Ticker
	OrderBook map[string]ws.OrderBookSnapshot
}
