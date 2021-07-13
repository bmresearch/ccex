package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	models "github.com/murlokito/ccex/models/rest"
)

type (
	FuturesClient struct {
		client *rest.Client
	}
)

func (f FuturesClient) GetFutures(futures *models.RequestForFutures) (*models.ResponseForFutures, error) {
	panic("implement me")
}

func (f FuturesClient) GetFuture(future *models.RequestForFuture) (*models.ResponseForFuture, error) {
	panic("implement me")
}

func (f FuturesClient) GetOpenInterest(stats *models.RequestForOpenInterest) (*models.ResponseForOpenInterest, error) {
	panic("implement me")
}

func (f FuturesClient) GetFundingRate(rates *models.RequestForFundingRates) (*models.ResponseForFundingRates, error) {
	panic("implement me")
}

func (f FuturesClient) GetIndexWeights(weights *models.RequestForIndexWeights) (*models.ResponseForIndexWeights, error) {
	panic("implement me")
}

func (f FuturesClient) GetExpiredFutures(futures *models.RequestForExpiredFutures) (*models.ResponseForExpiredFutures, error) {
	panic("implement me")
}

func (f FuturesClient) GetHistoricalIndex(req *models.RequestForHistoricalIndex) (*models.ResponseForHistoricalIndex, error) {
	panic("implement me")
}

func (f FuturesClient) PostFuturesAccountLeverageChange(change *models.RequestForFuturesAccountLeverageChange) (*models.ResponseForFuturesAccountLeverageChange, error) {
	panic("implement me")
}

// NewFuturesClient returns a new configured account client
func NewFuturesClient(client *rest.Client) (*FuturesClient, error) {
	return &FuturesClient{
		client: client,
	}, nil
}