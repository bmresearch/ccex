package exchange

import (
	"github.com/murlokito/ccex/models/rest"
)

type (
	// Futures interface specifies functionality to interact with a Futures API.
	Futures interface {
		/*
			GetFutures fetches all the available futures markets.
			In certain cases this method and the GetMarkets method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetFutures(futures *rest.RequestForFutures) (*rest.ResponseForFutures, error)

		/*
			GetFuture is used to fetch information related to the futures market specified by `future`.
			In certain cases this method and the GetMarket method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetFuture(future *rest.RequestForFuture) (*rest.ResponseForFuture, error)

		// GetOpenInterest fetches the stats associated with the future specified by `future`.
		GetOpenInterest(stats *rest.RequestForOpenInterest)(*rest.ResponseForFutureStats, error)

		// GetFundingRate fetches the funding rates for the future specified by `future`.
		GetFundingRate(rates *rest.RequestForFundingRates) (*rest.ResponseForFundingRates, error)

		// GetIndexWeights fetches the weights of the index specified by `index`.
		GetIndexWeights(weights *rest.RequestForIndexWeights) (*rest.ResponseForIndexWeights, error)

		// GetExpiredFutures fetches futures that have expired.
		GetExpiredFutures(futures *rest.RequestForExpiredFutures) (*rest.ResponseForExpiredFutures, error)

		/*
			GetHistoricalIndex fetches OHLC data for the index specified by `index`
			with a maximum specified by `limit` for the passed `symbol` with the specified `resolution`.
			Optionally provide `start` and `end` to request a specific period.
		*/
		GetHistoricalIndex(req *rest.RequestForHistoricalIndex) (*rest.ResponseForHistoricalIndex, error)
	}
)
