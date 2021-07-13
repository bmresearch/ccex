package rest

import "time"

type (

	RequestForWalletCoins struct {

	}

	WalletCoin struct {

	}

	ResponseForWalletCoins struct {
		BaseResponse

	}

	RequestForWalletBalances struct {

	}

	WalletBalance struct {

	}

	ResponseForWalletBalances struct {
		BaseResponse
		Result []WalletBalance `json:"result"`
	}

	RequestForDepositAddress struct {

	}

	DepositAddress struct {
		Coin string
		Method string
	}

	ResponseForDepositAddress struct {
		BaseResponse

	}

	RequestForWalletDepositHistory struct {
		Limit int
		Start time.Time
		End time.Time
	}

	WalletDepositHistory struct {

	}

	ResponseForWalletDepositHistory struct {
		BaseResponse
	}

	RequestForWalletWithdrawalHistory struct {
		Limit int
		Start time.Time
		End time.Time
	}

	WalletWithdrawalHistory struct {

	}

	ResponseForWalletWithdrawalHistory struct {
		BaseResponse

	}

	RequestForSavedAddresses struct {
		Coin string
	}

	SavedAddress struct {

	}

	ResponseForSavedAddresses struct {
		BaseResponse

	}

	RequestForSavedAddressCreation struct {
		Coin string
		Address string
		AddressName string
		Tag string
		PrimeTrust bool
	}

	ResponseForSavedAddressCreation struct {
		BaseResponse
	}

	RequestForSavedAddressDeletion struct {
		AddressID int
	}

	ResponseForSavedAddressDeletion struct {
		BaseResponse
	}

	RequestForWalletWithdrawal struct {
		Coin string
		Address string
		Tag string
		Password string
		Code string
		Size int
	}

	ResponseForWalletWithdrawal struct {
		BaseResponse
	}

)