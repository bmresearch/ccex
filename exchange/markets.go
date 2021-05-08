package exchange

import (
	"github.com/murlokito/ccex/models/rest"
)

type (
	Markets interface {
		/*
			GetMarkets fetches all the available markets.
			In certain cases this method and the GetFutures method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetMarkets(req *rest.RequestForMarkets) (*rest.ResponseForMarkets, error)

		/*
			GetMarket is used to fetch information related to the market specified by `Symbol`.
			In certain cases this method and the GetFuture method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetMarket(req *rest.RequestForMarket) (*rest.ResponseForMarket, error)

		// GetOrderBook fetches the order book for the passed `Market`.
		GetOrderBook(req *rest.RequestForOrderBook) (*rest.ResponseForOrderBook, error)

		// GetTrades fetches the trades for the passed `Symbol`.
		GetTrades(req *rest.RequestForTrades) (*rest.ResponseForTrades, error)

		/*
			GetCandles fetches OHLC data for the market specified by `Market`
			with a maximum specified by `Limit` for the passed `Market` with the specified `Resolution`.
			Optionally provide `Start` and `End` to request a specific period.
		*/
		GetCandles(req *rest.RequestForCandles) (*rest.ResponseForCandles, error)
	}
)
