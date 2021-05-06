package exchange

import "github.com/murlokito/ccex/common"

type (
	Conversion interface {
		PostConversionQuoteRequest(fromCoin, toCoin string, size float32) (common.Response, error)

		GetConversionQuoteStatus(quoteId int, market string) (common.Response, error)

		PostConversionQuoteAcceptance(quoteId int) (common.Response, error)
	}
)
