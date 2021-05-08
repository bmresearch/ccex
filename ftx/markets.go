package ftx

import (
	"github.com/murlokito/ccex/ftx/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (

	// MarketsClient represents the client for the FTX Markets API.
	MarketsClient struct {
		client *rest.Client
	}
)

func (m MarketsClient) GetMarkets(req *models.RequestForMarkets) (*models.ResponseForMarkets, error) {
	panic("implement me")
}

func (m MarketsClient) GetMarket(req *models.RequestForMarket) (*models.ResponseForMarket, error) {
	panic("implement me")
}

func (m MarketsClient) GetOrderBook(req *models.RequestForOrderBook) (*models.ResponseForOrderBook, error) {
	panic("implement me")
}

func (m MarketsClient) GetTrades(req *models.RequestForTrades) (*models.ResponseForTrades, error) {
	panic("implement me")
}

func (m MarketsClient) GetCandles(req *models.RequestForCandles) (*models.ResponseForCandles, error) {
	panic("implement me")
}

// NewMarketsClient returns a new configured account client
func NewMarketsClient(client *rest.Client) (*MarketsClient, error) {
	return &MarketsClient{
		client: client,
	}, nil
}

