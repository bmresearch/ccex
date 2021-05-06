package ftx

import "github.com/murlokito/ccex/ftx/rest"

type (

	// OrdersClient represents the client for the FTX Orders API.
	OrdersClient struct {
		client *rest.Client
	}
)

func NewOrdersClient(client *rest.Client) (*OrdersClient, error) {
	return &OrdersClient{
		client: client,
	}, nil
}
