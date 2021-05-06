package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
	"time"
)

type (

	// FillsClient represents the client for the FTX Fills API.
	FillsClient struct {
		client *rest.Client
	}
)

func (f FillsClient) GetFills(market, order string, orderId, limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

// NewFillsClient returns a new configured account client
func NewFillsClient(client *rest.Client) (*FillsClient, error) {
	return &FillsClient{
		client: client,
	}, nil
}

