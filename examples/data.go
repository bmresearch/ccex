package examples

import "github.com/murlokito/ccex/models/ws"

type Data struct {
	Trades map[string][]ws.TradeData
	Ticker map[string]ws.TickerData
	OrderBook map[string]ws.OrderBookData
	Markets ws.MarketData
}
