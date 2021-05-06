package exchange

import "github.com/murlokito/ccex/common"

type (
	Margin interface {
		GetLendingHistory() (common.Response, error)

		GetBorrowRates() (common.Response, error)

		GetLendingRates() (common.Response, error)

		GetDailyBorrowedAmounts() (common.Response, error)

		GetSpotMarginMarketInfo() (common.Response, error)

		GetMyBorrowHistory() (common.Response, error)

		GetMyLendingHistory() (common.Response, error)

		GetLendingOffers() (common.Response, error)

		GetLendingInfo() (common.Response, error)

		PostLendingOffer(coin string, size, rate float32) (common.Response, error)
	}
)
