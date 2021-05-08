package exchange

import (
	"github.com/murlokito/ccex/models/rest"
)

type (
	// Wallet specifies functionality for the wallet API
	Wallet interface {

		// GetWalletBalances is used to fetch balances of account holdings.
		GetWalletBalances(req *rest.RequestForWalletBalances) (rest.ResponseForWalletBalances, error)

		/*
			GetDepositAddress is used to fetch a deposit address for the currency specified by `coin` using `method`.
			For FTX, `method` can be one of:
				For ERC20 tokens: method=erc20
				For TRC20 tokens: method=trx
				For SPL tokens: method=sol
				For Omni tokens: method=omni
				For BEP2 tokens: method=bep2
		*/
		GetDepositAddress(req *rest.RequestForDepositAddress) (rest.ResponseForDepositAddress, error)

		// GetWalletDepositHistory is used to fetch the wallet deposit history.
		GetWalletDepositHistory(req *rest.RequestForWalletDepositHistory) (rest.ResponseForWalletDepositHistory, error)

		// GetWalletWithdrawalHistory is used to fetch the wallet withdrawal history.
		GetWalletWithdrawalHistory(req *rest.RequestForWalletWithdrawalHistory) (rest.ResponseForWalletWithdrawalHistory, error)

		// GetSavedAddresses is used to fetch saved addresses for currency specified by `coin`.
		GetSavedAddresses(req *rest.RequestForSavedAddresses) (rest.ResponseForSavedAddresses, error)

		// PostCreateSavedAddress is used to fetch saved addresses for currency specified by `coin`.
		PostCreateSavedAddress(req *rest.RequestForSavedAddressCreation) (rest.ResponseForSavedAddressCreation, error)

		// DeleteSavedAddress is used to delete a saved address specified by `addressId`.
		DeleteSavedAddress(req *rest.RequestForSavedAddressDeletion) (rest.ResponseForSavedAddressDeletion, error)

		/*
			PostWalletWithdrawal is used to request a withdrawal of the coin specified by `coin` with amount specified by `size`
			to the address specified by `address`, if withdrawal password and/or 2FA is active then `password` and `code`
			are necessary to proceed with the withdrawal. The parameter `tag` is optional.
		*/
		PostWalletWithdrawal(req *rest.RequestForWalletWithdrawal) (rest.ResponseForWalletWithdrawal, error)
	}
)
