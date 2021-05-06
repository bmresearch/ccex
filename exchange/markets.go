package exchange

import (
	"github.com/murlokito/ccex/common"
	"time"
)

type (
	Markets interface {
		/*
			GetMarkets fetches all the available markets.
			In certain cases this method and the GetFutures method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetMarkets() (common.Response, error)

		/*
			GetMarket is used to fetch information related to the market specified by `symbol`.
			In certain cases this method and the GetFuture method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetMarket(symbol string) (common.Response, error)

		// GetOrderBook fetches the order book for the passed `symbol`.
		GetOrderBook(symbol string) (common.Response, error)

		// GetTrades fetches the trades for the passed `symbol`.
		GetTrades(symbol string) (common.Response, error)

		/*
			GetCandles fetches OHLC data for the market specified by `symbol`
			with a maximum specified by `limit` for the passed `symbol` with the specified `resolution`.
			Optionally provide `start` and `end` to request a specific period.
		*/
		GetCandles(symbol string, resolution, limit int, start, end time.Time) (common.Response, error)
	}
)
