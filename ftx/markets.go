package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
	"time"
)

type (

	// MarketsClient represents the client for the FTX Markets API.
	MarketsClient struct {
		client *rest.Client
	}
)

func (m MarketsClient) GetMarkets() (common.Response, error) {
	panic("implement me")
}

func (m MarketsClient) GetMarket(symbol string) (common.Response, error) {
	panic("implement me")
}

func (m MarketsClient) GetOrderBook(symbol string) (common.Response, error) {
	panic("implement me")
}

func (m MarketsClient) GetTrades(symbol string) (common.Response, error) {
	panic("implement me")
}

func (m MarketsClient) GetCandles(symbol string, resolution, limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

// NewMarketsClient returns a new configured account client
func NewMarketsClient(client *rest.Client) (*MarketsClient, error) {
	return &MarketsClient{
		client: client,
	}, nil
}

