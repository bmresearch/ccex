package rest

import "time"

type (

	// RequestForFutures represents a request for a futures data
	RequestForFutures struct {
	}

	// RequestForFuture represents a request for a future's data
	RequestForFuture struct {
		Future string
	}

	// Future holds the necessary information to represent a future
	Future struct {
		Ask                 float64   `json:"ask"`
		Bid                 float64   `json:"bid"`
		Change1H            float64   `json:"change1h"`
		Change24H           float64   `json:"change24h"`
		ChangeBod           float64   `json:"changeBod,omitempty"`
		VolumeUsd24H        float64   `json:"volumeUsd24h,omitempty"`
		Volume              float64   `json:"volume,omitempty"`
		Description         string    `json:"description"`
		Enabled             bool      `json:"enabled"`
		Expired             bool      `json:"expired"`
		Expiry              time.Time `json:"expiry"`
		Index               float64   `json:"index"`
		ImfFactor           float64   `json:"imfFactor"`
		Last                float64   `json:"last"`
		LowerBound          float64   `json:"lowerBound"`
		Mark                float64   `json:"mark"`
		Name                string    `json:"name"`
		Perpetual           bool      `json:"perpetual"`
		PositionLimitWeight float64   `json:"positionLimitWeight,omitempty"`
		PostOnly            bool      `json:"postOnly"`
		PriceIncrement      float64   `json:"priceIncrement"`
		SizeIncrement       float64   `json:"sizeIncrement"`
		Underlying          string    `json:"underlying"`
		UpperBound          float64   `json:"upperBound"`
		Type                string    `json:"type"`
	}

	// ResponseForFutures holds the necessary information to represent the response for futures data
	ResponseForFutures struct {
		BaseResponse
		Result []Future `json:"result"`
	}

	// ResponseForFuture holds the necessary information to represent the response for a future's data
	ResponseForFuture struct {
		BaseResponse
		Result Future `json:"result"`
	}

	// FutureStats holds the necessary information to represent the future's statistics
	// TODO: Remove and substitute with an open interest specific
	FutureStats struct {
		Volume                   float64   `json:"volume"`
		NextFundingRate          float64   `json:"nextFundingRate"`
		NextFundingTime          time.Time `json:"nextFundingTime"`
		ExpirationPrice          float64   `json:"expirationPrice"`
		PredictedExpirationPrice float64   `json:"predictedExpirationPrice"`
		StrikePrice              float64   `json:"strikePrice"`
		OpenInterest             float64   `json:"openInterest"`
	}

	// ResponseForFutureStats holds the necessary information to represent the response for a future's dat
	// TODO: Remove and substitute with an open interest specific
	ResponseForFutureStats struct {
		BaseResponse
		Result FutureStats `json:"result"`
	}

	// RequestForFundingRates represents a request for funding rates data.
	RequestForFundingRates struct {
		Future string
		Start  time.Time
		End    time.Time
	}

	// FundingRate holds the necessary information to represent the funding rate
	FundingRate struct {
		Future string    `json:"future"`
		Rate   float64   `json:"rate"`
		Time   time.Time `json:"time"`
	}

	// ResponseForFundingRates holds the necessary information to represent the response for a future's data
	ResponseForFundingRates struct {
		BaseResponse
		Result []FundingRate `json:"result"`
	}

	// FundingPayment holds the necessary information to represent a funding payment
	FundingPayment struct {
		Future  string    `json:"future"`
		ID      int64     `json:"id"`
		Payment float64   `json:"payment"`
		Rate    float64   `json:"rate"`
		Time    time.Time `json:"time"`
	}

	// RequestForFundingPayments ..
	// TODO: to keep or not to keep? tune in at 5
	RequestForFundingPayments struct {
		Future string
		Start  time.Time
		End    time.Time
	}

	// ResponseForFundingPayments holds the necessary information to represent the response for a future's data
	ResponseForFundingPayments struct {
		BaseResponse
		Result []FundingPayment `json:"result"`
	}

 	// RequestForIndexWeights represents a request for index weights data
	RequestForIndexWeights struct {
		Index string
	}

	// IndexWeight holds the necessary...
	IndexWeight struct {
	}

	// RequestForHistoricalIndex represents a request for historical index data
	RequestForHistoricalIndex struct {
		Index      string
		Resolution int
		Limit      int
		Start      time.Time
		End        time.Time
	}

	// HistoricalIndex holds the necessary...
	HistoricalIndex struct {
	}

	RequestForOpenInterest struct {
		Future string
	}

	RequestForExpiredFutures struct {
	}
)
