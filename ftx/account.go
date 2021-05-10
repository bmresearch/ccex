package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
	// AccountClient represents the client for the FTX Account API.
	AccountClient struct {
		client *rest.Client
	}
)

func (a AccountClient) GetPositions(req *models.RequestForPositions) (*models.ResponseForPositions, error) {
	panic("implement me")
}

func (a AccountClient) PostAccountLeverageChange(req *models.RequestForAccountLeverageChange) (*models.ResponseForAccountLeverageChange, error) {
	panic("implement me")
}


// NewAccountClient returns a new configured account client
func NewAccountClient(client *rest.Client) (*AccountClient, error) {
	return &AccountClient{
		client: client,
	}, nil
}
