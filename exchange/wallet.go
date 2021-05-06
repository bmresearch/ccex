package exchange

import (
	"github.com/murlokito/ccex/common"
	"time"
)

type (
	// Wallet specifies functionality for the wallet API
	Wallet interface {

		// GetWalletCoins is used to fetch coins held in the account wallet.
		GetWalletCoins() (common.Response, error)

		// GetWalletBalances is used to fetch balances of account holdings.
		GetWalletBalances() (common.Response, error)

		// GetAllWalletBalances is used to fetch all balances of account holdings.
		GetAllWalletBalances() (common.Response, error)

		/*
			GetDepositAddress is used to fetch a deposit address for the currency specified by `coin` using `method`.
			For FTX, `method` can be one of:
				For ERC20 tokens: method=erc20
				For TRC20 tokens: method=trx
				For SPL tokens: method=sol
				For Omni tokens: method=omni
				For BEP2 tokens: method=bep2
		*/
		GetDepositAddress(coin, method string) (common.Response, error)

		// GetWalletDepositHistory is used to fetch the wallet deposit history.
		GetWalletDepositHistory(limit int, start, end time.Time) (common.Response, error)

		// GetWalletWithdrawalHistory is used to fetch the wallet withdrawal history.
		GetWalletWithdrawalHistory(limit int, start, end time.Time) (common.Response, error)

		// GetWalletAirdropHistory is used to fetch the wallet airdrop history.
		GetWalletAirdropHistory(limit int, start, end time.Time) (common.Response, error)

		// GetSavedAddresses is used to fetch saved addresses for currency specified by `coin`.
		GetSavedAddresses(coin string) (common.Response, error)

		// PostCreateSavedAddress is used to fetch saved addresses for currency specified by `coin`.
		PostCreateSavedAddress(coin, address, addressName, tag string, isPrimeTrust bool) (common.Response, error)

		// DeleteSavedAddress is used to delete a saved address specified by `addressId`.
		DeleteSavedAddress(addressId int) (common.Response, error)

		/*
			PostWalletWithdrawal is used to request a withdrawal of the coin specified by `coin` with amount specified by `size`
			to the address specified by `address`, if withdrawal password and/or 2FA is active then `password` and `code`
			are necessary to proceed with the withdrawal. The parameter `tag` is optional.
		*/
		PostWalletWithdrawal(coin, address, tag, password, code string, size int) (common.Response, error)
	}
)
