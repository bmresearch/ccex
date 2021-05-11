package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
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