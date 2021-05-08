package rest

type (

	// ResponseForAccount holds the necessary information to represent the response the account data
	ResponseForAccount struct {
		BaseResponse
		Result Account `json:"result"`
	}




	// ResponseForIndexWeights holds the necessary information to represent the response for index weights
	ResponseForIndexWeights struct {
		BaseResponse
		Result []IndexWeight `json:"result"`
	}

	ResponseForHistoricalIndex struct {
		BaseResponse
		Result HistoricalIndex `json:"result"`
	}

	ResponseForExpiredFutures struct {
		BaseResponse
		Result []Future `json:"result"`
	}
)
