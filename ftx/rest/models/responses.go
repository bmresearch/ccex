package models

/*
Account requests responses
*/

// ResponseForAccount holds the necessary information to represent the response the account data
type ResponseForAccount struct {
	BaseResponse
	Result Account `json:"result"`
}

func (r *ResponseForAccount) GetResult() interface{} {
	return r.Result
}

// ResponseForSubAccounts holds the necessary information to represent the response for sub account data
type ResponseForSubAccounts struct {
	BaseResponse
	Result []SubAccount `json:"result"`
}

func (r *ResponseForSubAccounts) GetResult() interface{} {
	return r.Result
}

// ResponseForSubAccountBalance holds the necessary information to represent the response for sub account balances
type ResponseForSubAccountBalance struct {
	BaseResponse
	Result []SubAccountBalance `json:"result"`
}

func (r *ResponseForSubAccountBalance) GetResult() interface{} {
	return r.Result
}

// ResponseForSubAccountCreation holds the necessary information to represent the response for a sub account creation
type ResponseForSubAccountCreation struct {
	BaseResponse
	Result SubAccount `json:"result"`
}

func (r *ResponseForSubAccountCreation) GetResult() interface{} {
	return r.Result
}

// ResponseForSubAccountTransfer holds the necessary information to represent the response for a sub account transfer
type ResponseForSubAccountTransfer struct {
	BaseResponse
	Result SubAccountTransfer `json:"result"`
}

func (r *ResponseForSubAccountTransfer) GetResult() interface{} {
	return r.Result
}

// ResponseForPositions holds the necessary information to represent the response for positions data
type ResponseForPositions struct {
	BaseResponse
	Result []Position `json:"result"`
}

func (r *ResponseForPositions) GetResult() interface{} {
	return r.Result
}

/*
Wallet requests responses
*/




/*
Market requests responses
*/

// ResponseForCandles holds the necessary information to represent the response for candlestick data
type ResponseForCandles struct {
	BaseResponse
	Result []Candle `json:"result"`
}

func (r *ResponseForCandles) GetResult() interface{} {
	return r.Result
}

// ResponseForMarkets holds the necessary information to represent the response for markets data
type ResponseForMarkets struct {
	BaseResponse
	Result []Market `json:"result"`
}

func (r *ResponseForMarkets) GetResult() interface{} {
	return r.Result
}

// ResponseForMarket holds the necessary information to represent the response for a market's data
type ResponseForMarket struct {
	BaseResponse
	Result Market `json:"result"`
}

func (r *ResponseForMarket) GetResult() interface{} {
	return r.Result
}

// ResponseForFutures holds the necessary information to represent the response for futures data
type ResponseForFutures struct {
	BaseResponse
	Result []Future `json:"result"`
}

func (r *ResponseForFutures) GetResult() interface{} {
	return r.Result
}

// ResponseForFuture holds the necessary information to represent the response for a future's data
type ResponseForFuture struct {
	BaseResponse
	Result Future `json:"result"`
}

func (r *ResponseForFuture) GetResult() interface{} {
	return r.Result
}

// ResponseForFutureStats holds the necessary information to represent the response for a future's data
type ResponseForFutureStats struct {
	BaseResponse
	Result FutureStats `json:"result"`
}

func (r *ResponseForFutureStats) GetResult() interface{} {
	return r.Result
}

// ResponseForFundingRates holds the necessary information to represent the response for a future's data
type ResponseForFundingRates struct {
	BaseResponse
	Result []FundingRate `json:"result"`
}

func (r *ResponseForFundingRates) GetResult() interface{} {
	return r.Result
}

// ResponseForOrderBook holds the necessary information to represent the response for order book data
type ResponseForOrderBook struct {
	BaseResponse
	Result OrderBook `json:"result"`
}

func (r *ResponseForOrderBook) GetResult() interface{} {
	return r.Result
}

// ResponseForTrades holds the necessary information to represent the response for trade data
type ResponseForTrades struct {
	BaseResponse
	Result []Trade `json:"result"`
}

func (r *ResponseForTrades) GetResult() interface{} {
	return r.Result
}



