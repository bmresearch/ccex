package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
	"time"
)

type (

	// FuturesClient represents the client for the FTX Futures API.
	FuturesClient struct {
		client *rest.Client
	}
)

func (f FuturesClient) GetFutures() (common.Response, error) {
	panic("implement me")
}

func (f FuturesClient) GetFuture(future string) (common.Response, error) {
	panic("implement me")
}

func (f FuturesClient) GetFutureStats(future string) (common.Response, error) {
	panic("implement me")
}

func (f FuturesClient) GetFundingRate(future string, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

func (f FuturesClient) GetIndexWeights(index string) (common.Response, error) {
	panic("implement me")
}

func (f FuturesClient) GetExpiredFutures() (common.Response, error) {
	panic("implement me")
}

func (f FuturesClient) GetHistoricalIndex(index string, resolution, limit int, start, end time.Time) (common.Response, error) {
	panic("implement me")
}

// NewFuturesClient returns a new configured account client
func NewFuturesClient(client *rest.Client) (*FuturesClient, error) {
	return &FuturesClient{
		client: client,
	}, nil
}

