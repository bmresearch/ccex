package exchange

import (
	"github.com/murlokito/ccex/models/rest"
)

type (
	// SubAccount interface specifies functionality to interact with a SubAccount API.
	SubAccount interface {

		// GetSubAccounts is used to get all sub accounts.
		GetSubAccounts(req *rest.RequestForSubAccounts) (rest.ResponseForSubAccounts, error)

		// GetSubAccountBalance is used to get the balance of the sub account specified by `Name`.
		GetSubAccountBalance(req *rest.RequestForSubAccountBalance) (rest.ResponseForSubAccountBalance, error)

		/*
			PostSubAccountNameChange is used to change the name of the sub account
			specified by `Name` to a new name specified by `New`.
		*/
		PostSubAccountNameChange(req *rest.RequestForSubAccountChange) (rest.ResponseForSubAccountChange, error)

		/*
			PostSubAccountTransfer is used to transfer the asset specified by `Coin`
			from `Source` to a sub account specified by `Target` with the amount of `Size`.
		*/
		PostSubAccountTransfer(req *rest.RequestForSubAccountTransfer) (rest.ResponseForSubAccountTransfer, error)
	}
)
