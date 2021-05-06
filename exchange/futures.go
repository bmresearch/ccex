package exchange

import (
	"github.com/murlokito/ccex/common"
	"time"
)

type (
	Futures interface {
		/*
			GetFutures fetches all the available futures markets.
			In certain cases this method and the GetMarkets method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetFutures() (common.Response, error)

		/*
			GetFuture is used to fetch information related to the futures market specified by `future`.
			In certain cases this method and the GetMarket method have the same behavior,
			it happens when it is strictly a futures exchange or there are no distinct endpoints for markets and futures.
			It is so the interface is implemented.
		*/
		GetFuture(future string) (common.Response, error)

		// GetFutureStats fetches the stats associated with the future specified by `future`.
		GetFutureStats(future string) (common.Response, error)

		// GetFundingRate fetches the funding rates for the future specified by `future`.
		GetFundingRate(future string, start, end time.Time) (common.Response, error)

		// GetIndexWeights fetches the weights of the index specified by `index`.
		GetIndexWeights(index string) (common.Response, error)

		// GetExpiredFutures fetches futures that have expired.
		GetExpiredFutures() (common.Response, error)

		/*
			GetHistoricalIndex fetches OHLC data for the index specified by `index`
			with a maximum specified by `limit` for the passed `symbol` with the specified `resolution`.
			Optionally provide `start` and `end` to request a specific period.
		*/
		GetHistoricalIndex(index string, resolution, limit int, start, end time.Time) (common.Response, error)
	}
)
