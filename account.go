package ccex

import "github.com/murlokito/ccex/common"

type (
	// Account interface specifies the functionality for the account API
	Account interface {
		// GetAccount is used to get information associated with the account.
		GetAccount() (common.Response, error)

		// GetPositions is used to get the open positions.
		GetPositions(showAvgPrice bool) (common.Response, error)

		// GetSubAccounts is used to get all sub accounts.
		GetSubAccounts() (common.Response, error)

		// GetSubAccountBalance is used to get the balance of the sub account specified by `subAccount`.
		GetSubAccountBalance(subAccount string) (common.Response, error)

		/*
			PostSubAccountNameChange is used to change the name of the sub account
			specified by `subAccountName` to a new name specified by `newSubAccountName`.
		*/
		PostSubAccountNameChange(subAccountName, newSubAccountName string) (common.Response, error)

		/*
			PostSubAccountAssetTransfer is used to transfer the asset specified by `coin`
			from `sourceSubAccount` to a sub account specified by `targetSubAccount` with the amount of `size`.
		*/
		PostSubAccountAssetTransfer(coin, sourceSubAccount, targetSubAccount string, size int) (common.Response, error)

		// PostAccountLeverageChange is used to change the account's maximum leverage to the amount specified by `leverage`.
		PostAccountLeverageChange(leverage int) (common.Response, error)

		// PostFuturesAccountLeverageChange is used to change the futures account's maximum leverage to the amount specified by `leverage` on the market specified by `symbol`.
		PostFuturesAccountLeverageChange(symbol string, leverage int) (common.Response, error)

	}
)
