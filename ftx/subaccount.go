package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
	// SubAccountClient  represents the client for the FTX SubAccount API.
	SubAccountClient struct {
		client *rest.Client
	}
)

// GetSubAccounts fetches the sub accounts.
func (a SubAccountClient) GetSubAccounts(req *models.RequestForSubAccounts) (models.ResponseForSubAccounts, error) {
	panic("implement me")
}

// GetSubAccountBalance fetches the sub accounts balance.
func (a SubAccountClient) GetSubAccountBalance(req *models.RequestForSubAccountBalance) (models.ResponseForSubAccountBalance, error) {
	panic("implement me")
}

// PostSubAccountNameChange requests a sub account name change.
func (a SubAccountClient) PostSubAccountNameChange(req *models.RequestForSubAccountChange) (models.ResponseForSubAccountChange, error) {
	panic("implement me")
}

// PostSubAccountTransfer requests a sub account asset transfer.
func (a SubAccountClient) PostSubAccountTransfer(req *models.RequestForSubAccountTransfer) (models.ResponseForSubAccountTransfer, error) {
	panic("implement me")
}

// NewSubAccountClient returns a new configured sub account client.
func NewSubAccountClient(client *rest.Client) (*SubAccountClient, error) {
	return &SubAccountClient{
		client: client,
	}, nil
}