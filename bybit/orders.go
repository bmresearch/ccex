package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
	OrdersClient struct {
		client *rest.Client
	}
)

func (o OrdersClient) GetOpenOrders(req *models.RequestForOpenOrders) (models.ResponseForOpenOrders, error) {
	panic("implement me")
}

func (o OrdersClient) GetOrderHistory(req *models.RequestForOrderHistory) (models.ResponseForOrderHistory, error) {
	panic("implement me")
}

// NewOrdersClient returns a new configured account client
func NewOrdersClient(client *rest.Client) (*OrdersClient, error) {
	return &OrdersClient{
		client: client,
	}, nil
}