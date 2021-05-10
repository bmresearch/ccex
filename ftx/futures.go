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
	url := fmt.Sprintf(rest.FuturesEndpointFormat, req.Future)

	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}
	var model models.ResponseForFuture

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

func (f FuturesClient) GetOpenInterest(req *models.RequestForOpenInterest) (*models.ResponseForOpenInterest, error) {
	if req.Future == "" {
		return nil, fmt.Errorf("invalid params, must pass `Future`")
	}

	url := fmt.Sprintf(rest.FuturesStatsFormat, req.Future)

	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}
	var model models.ResponseForOpenInterest

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
	if req.Index == "" {
		return nil, fmt.Errorf("invalid params, must pass `Index`")
	}

	url := fmt.Sprintf(rest.IndexWeightsFormat, req.Index)


	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}

	var model models.ResponseForIndexWeights

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

func (f FuturesClient) GetExpiredFutures(req *models.RequestForExpiredFutures) (*models.ResponseForExpiredFutures, error) {
	url := rest.ExpiredFuturesEndpoint

	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}
	var model models.ResponseForExpiredFutures

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

func (f FuturesClient) GetHistoricalIndex(req *models.RequestForHistoricalIndex) (*models.ResponseForHistoricalIndex, error) {
	if req.Index == "" {
		return nil, fmt.Errorf("invalid params, must pass `Index`")
	}
	if req.Resolution == 0 {
		return nil, fmt.Errorf("invalid params, must pass `Resolution`")
	}

	var url string

	if req.Limit != 0 {
		if (req.Start != time.Time{}) && (req.End != time.Time{}) {
			url = fmt.Sprintf(rest.IndexCandlesLimitStartEndEndpointFormat, req.Index, req.Resolution, req.Limit, req.Start.Unix(), req.End.Unix())
		} else {
			url = fmt.Sprintf(rest.IndexCandlesLimitEndpointFormat, req.Index, req.Resolution, req.Limit)
		}
	} else {
		url = fmt.Sprintf(rest.IndexCandlesEndpointFormat, req.Index, req.Resolution)
	}

	res, err := f.client.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "error during request")
	}

	var model models.ResponseForHistoricalIndex

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

func (f FuturesClient) PostFuturesAccountLeverageChange(req *models.RequestForFuturesAccountLeverageChange) (*models.ResponseForFuturesAccountLeverageChange, error) {
	panic("implement me")
}

// NewFuturesClient returns a new configured account client
func NewFuturesClient(client *rest.Client) (*FuturesClient, error) {
	return &FuturesClient{
		client: client,
	}, nil
}

