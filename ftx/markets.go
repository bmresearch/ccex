package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	rest3 "github.com/murlokito/ccex/models/rest"
)

type (

	// MarketsClient represents the client for the FTX Markets API.
	MarketsClient struct {
		client *rest.Client
	}
)

func (m MarketsClient) GetMarkets(req *rest3.RequestForMarkets) (*rest3.ResponseForMarkets, error) {
	panic("implement me")
}

func (m MarketsClient) GetMarket(req *rest3.RequestForMarket) (*rest3.ResponseForMarket, error) {
	panic("implement me")
}

func (m MarketsClient) GetOrderBook(req *rest3.RequestForOrderBook) (*rest3.ResponseForOrderBook, error) {
	panic("implement me")
}

func (m MarketsClient) GetTrades(req *rest3.RequestForTrades) (*rest3.ResponseForTrades, error) {
	panic("implement me")
}

func (m MarketsClient) GetCandles(req *rest3.RequestForCandles) (*rest3.ResponseForCandles, error) {
	panic("implement me")
}

// NewMarketsClient returns a new configured account client
func NewMarketsClient(client *rest.Client) (*MarketsClient, error) {
	return &MarketsClient{
		client: client,
	}, nil
}

