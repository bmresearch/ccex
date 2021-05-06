package ftx

import (
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
)

type (

	// MarginClient represents the client for the FTX Margin API.
	MarginClient struct {
		client *rest.Client
	}
)

func (m MarginClient) GetLendingHistory() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetBorrowRates() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetLendingRates() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetDailyBorrowedAmounts() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetSpotMarginMarketInfo() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetMyBorrowHistory() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetMyLendingHistory() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetLendingOffers() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) GetLendingInfo() (common.Response, error) {
	panic("implement me")
}

func (m MarginClient) PostLendingOffer(coin string, size, rate float32) (common.Response, error) {
	panic("implement me")
}

// NewMarginClient returns a new configured account client
func NewMarginClient(client *rest.Client) (*MarginClient, error) {
	return &MarginClient{
		client: client,
	}, nil
}

