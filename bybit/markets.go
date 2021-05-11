package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
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