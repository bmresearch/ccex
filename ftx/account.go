package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	rest3 "github.com/murlokito/ccex/models/rest"
)

type (
	// AccountClient represents the client for the FTX Account API.
	AccountClient struct {
		client *rest.Client
	}
)

func (a AccountClient) GetPositions(req *rest3.RequestForPositions) (*rest3.ResponseForPositions, error) {
	panic("implement me")
}

func (a AccountClient) PostAccountLeverageChange(req *rest3.RequestForAccountLeverageChange) (*rest3.ResponseForAccountLeverageChange, error) {
	panic("implement me")
}


// NewAccountClient returns a new configured account client
func NewAccountClient(client *rest.Client) (*AccountClient, error) {
	return &AccountClient{
		client: client,
	}, nil
}
