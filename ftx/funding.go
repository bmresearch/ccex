package ftx

import (
	"encoding/json"
	"fmt"
	"github.com/murlokito/ccex/common"
	"github.com/murlokito/ccex/ftx/rest"
	"github.com/murlokito/ccex/ftx/rest/models"
	"time"
)

type (

	// FundingClient represents the client for the FTX Funding API.
	FundingClient struct {
		client *rest.Client
	}
)

func (f FundingClient) GetFundingPayments(future string, start, end time.Time) (common.Response, error) {
	var url string

	if future != "" {
		if (start != time.Time{}) && (end != time.Time{}) {
			url = fmt.Sprintf(rest.FundingPaymentsFutureStartEndEndpoint, future, start.Unix(), end.Unix())
		}else {
			url = fmt.Sprintf(rest.FundingPaymentsFutureEndpoint, future)
		}
	}

	res, err := f.client.Get(url)
	if err != nil {
		return nil, err
	}
	var model models.ResponseForFundingRates
	err = json.Unmarshal(res, &model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// NewFundingClient returns a new configured account client
func NewFundingClient(client *rest.Client) (*FundingClient, error) {
	return &FundingClient{
		client: client,
	}, nil
}

