package exchange

import (
	"github.com/murlokito/ccex/models/rest"
)

type (
	// Account interface specifies the functionality for the account API
	Account interface {

		// GetPositions is used to get the open positions.
		GetPositions(req *rest.RequestForPositions) (*rest.ResponseForPositions, error)

		// PostAccountLeverageChange is used to change the account's maximum leverage to the amount specified by `leverage`.
		PostAccountLeverageChange(req *rest.RequestForAccountLeverageChange) (*rest.ResponseForAccountLeverageChange, error)

		// PostFuturesAccountLeverageChange is used to change the futures account's maximum leverage to the amount specified by `leverage` on the market specified by `symbol`.
		PostFuturesAccountLeverageChange(change *rest.RequestForFuturesAccountLeverageChange) (*rest.ResponseForFuturesAccountLeverageChange, error)
	}
)
