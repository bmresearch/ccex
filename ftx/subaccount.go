package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	rest3 "github.com/murlokito/ccex/models/rest"
)

type (
	// SubAccountClient  represents the client for the FTX SubAccount API.
	SubAccountClient struct {
		client *rest.Client
	}
)

// GetSubAccounts fetches the sub accounts.
func (a SubAccountClient) GetSubAccounts(req *rest3.RequestForSubAccounts) (rest3.ResponseForSubAccounts, error) {
	panic("implement me")
}

// GetSubAccountBalance fetches the sub accounts balance.
func (a SubAccountClient) GetSubAccountBalance(req *rest3.RequestForSubAccountBalance) (rest3.ResponseForSubAccountBalance, error) {
	panic("implement me")
}

// PostSubAccountNameChange requests a sub account name change.
func (a SubAccountClient) PostSubAccountNameChange(req *rest3.RequestForSubAccountChange) (rest3.ResponseForSubAccountChange, error) {
	panic("implement me")
}

// PostSubAccountTransfer requests a sub account asset transfer.
func (a SubAccountClient) PostSubAccountTransfer(req *rest3.RequestForSubAccountTransfer) (rest3.ResponseForSubAccountTransfer, error) {
	panic("implement me")
}

// NewSubAccountClient returns a new configured sub account client.
func NewSubAccountClient(client *rest.Client) (*SubAccountClient, error) {
	return &SubAccountClient{
		client: client,
	}, nil
}