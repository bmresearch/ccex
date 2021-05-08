package ftx

import (
	"encoding/json"
	"fmt"
	"github.com/murlokito/ccex/ftx/rest"
	models "github.com/murlokito/ccex/models/rest"
	"github.com/pkg/errors"
	"time"
)

type (

	// FuturesClient represents the client for the FTX Futures API.
	FuturesClient struct {
		client *rest.Client
	}
)

func (f FuturesClient) GetFutures(req *models.RequestForFutures) (*models.ResponseForFutures, error) {
	url := rest.FuturesEndpoint

	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}
	var model models.ResponseForFutures

	if len(res) != 0 {
		err = json.Unmarshal(res, &model)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling request response")
		}
	} else {
		return nil, fmt.Errorf("something went wrong during request")
	}

	return &model, nil
}

func (f FuturesClient) GetFuture(req *models.RequestForFuture) (*models.ResponseForFuture, error) {
	panic("implement me")
}

func (f FuturesClient) GetOpenInterest(req *models.RequestForOpenInterest) (*models.ResponseForFutureStats, error) {
	panic("implement me")
}

func (f FuturesClient) GetFundingPayments(future string, start, end time.Time) (*models.ResponseForFundingPayments, error) {
	var url string

	if future != "" {
		if (start != time.Time{}) && (end != time.Time{}) {
			url = fmt.Sprintf(rest.FundingPaymentsFutureStartEndEndpoint, future, start.Unix(), end.Unix())
		}else {
			url = fmt.Sprintf(rest.FundingPaymentsFutureEndpoint, future)
		}
	} else {
		url = rest.FundingPaymentsEndpoint
	}

	res, err := f.client.Get(url)
	if err != nil {
		return nil, err
	}
	var model models.ResponseForFundingPayments
	err = json.Unmarshal(res, &model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (f FuturesClient) GetFundingRate(req *models.RequestForFundingRates) (*models.ResponseForFundingRates, error) {
	var url string

	if req.Future != "" {
		if (req.Start != time.Time{}) && (req.End != time.Time{}) {
			url = fmt.Sprintf(rest.FuturesFutureStartEndFundingRate, req.Future, req.Start.Unix(), req.End.Unix())
		} else {
			url = fmt.Sprintf(rest.FuturesFutureFundingRate, req.Future)
		}
	} else {
		url = rest.FuturesFundingRate
	}

	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}

	var model models.ResponseForFundingRates

	if len(res) != 0 {
		err = json.Unmarshal(res, &model)
		if err != nil {
			return nil, errors.Wrap(err, "error unmarshalling request response")
		}
	} else {
		return nil, fmt.Errorf("something went wrong during request")
	}

	return &model, nil
}

func (f FuturesClient) GetIndexWeights(req *models.RequestForIndexWeights) (*models.ResponseForIndexWeights, error) {
	panic("implement me")
}

func (f FuturesClient) GetExpiredFutures(req *models.RequestForExpiredFutures) (*models.ResponseForExpiredFutures, error) {
	panic("implement me")
}

func (f FuturesClient) GetHistoricalIndex(req *models.RequestForHistoricalIndex) (*models.ResponseForHistoricalIndex, error) {
	panic("implement me")
}

// NewFuturesClient returns a new configured account client
func NewFuturesClient(client *rest.Client) (*FuturesClient, error) {
	return &FuturesClient{
		client: client,
	}, nil
}

