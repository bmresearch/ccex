package rest

/*
	General constants needed to perform requests to the API.
*/
const (
	Url = "https://ftx.com/api"
)

/*
	The following constants define the API endpoints for FTX.
	There are simple endpoint definitions which represent a string constant
	and more complex endpoint definitions that represent a string format
	which is to be used by an helper method to format the string with certain parameters.
*/
const (
	/*
		Account Endpoints
	*/

	// The AccountEndpoint is used as GET to fetch the account information
	AccountEndpoint = "/account"
	// The AccountLeverageChangeEndpoint is used as POST to change the maximum account leverage
	AccountLeverageChangeEndpoint = "/account/leverage"
	// The PositionsEndpoint is used as GET to fetch the open positions
	PositionsEndpoint = "/positions"
	// The SubAccountsEndpoint is used as GET to fetch all sub accounts
	SubAccountsEndpoint = "/subaccounts"
	// The SubAccountChangeNameEndpoint is used as POST to change a sub account name
	SubAccountChangeNameEndpoint = "/subaccounts/update_name"
	// The SubAccountTransferEndpoint is used as POST to transfer assets between sub accounts
	SubAccountTransferEndpoint = "/subaccounts/transfer"

	/*
		Account Endpoint Formats
	*/

	// The SubAccountBalanceEndpointFormat is used as GET to fetch a sub account balance
	SubAccountBalanceEndpointFormat = "/subaccounts/%s/balance"

	/*
		Wallet Endpoints
	*/

	// The WalletCoinsEndpoint is used as GET to fetch the wallet's holding coins
	WalletCoinsEndpoint = "/wallet/coins"
	// The WalletBalancesEndpoint is used as GET to fetch the wallet balances
	WalletBalancesEndpoint = "/wallet/balances"
	// The WalletAllBalancesEndpoint is used as GET to fetch all wallet balances
	WalletAllBalancesEndpoint = "/wallet/balances"
	// The WalletDepositHistoryEndpoint is used as GET to fetch the wallet deposit history
	WalletDepositHistoryEndpoint = "/wallet/deposits"
	// The WalletWithdrawalEndpoint is used as GET to fetch withdrawal history and POST to request new withdrawal
	WalletWithdrawalEndpoint = "/wallet/withdrawals"
	// The WalletAirdropsEndpoint is used as GET to fetch the wallet's airdrop history
	WalletAirdropsEndpoint = "/wallet/airdrops"
	// The WalletSavedAddressesEndpoint is used as GET to fetch saved addresses and POST to save a new one
	WalletSavedAddressesEndpoint = "/wallet/saved_addresses"

	/*
		Wallet Endpoint Formats
	*/

	// The WalletDepositAddressEndpointFormat
	WalletDepositAddressEndpointFormat = "/wallet/deposit_address/%s?method=%s"
	// The WalletSavedAddressDeleteEndpointFormat
	WalletSavedAddressDeleteEndpointFormat = "/wallet/saved_addresses/%s"

	/*
		Markets Endpoints
	*/

	// The MarketsEndpoint is used as GET to fetch all the available markets
	MarketsEndpoint = "/markets"

	/*
		Markets Endpoint Formats
	*/

	// The MarketsEndpointFormat
	MarketsEndpointFormat = "/markets/%s/%s"
	// The MarketsOrderBookEndpointFormat
	MarketsOrderBookEndpointFormat = "/markets/%s/%s/orderbook?depth=%s"
	// The TradesEndpointFormat
	TradesEndpointFormat = "/markets/%s/%s/trades?limit=%s&start_time=%s&end_time=%s"
	// The CandlesEndpointFormat is used as GET to fetch candlesticks for a certain market with the default limit of 1500
	CandlesEndpointFormat = "/markets/%s/%s/candles?resolution=%s"
	// The CandlesLimitEndpointFormat is used as GET to fetch candlesticks for a certain market with a specified limit
	CandlesLimitEndpointFormat = "/markets/%s/%s/candles?resolution=%s&limit=%s"
	// The CandlesLimitFromEndpointFormat is used as GET to fetch candlesticks for a certain market with a specified limit from a certain time
	CandlesLimitFromEndpointFormat = "/markets/%s/%s/candles?resolution=%s&limit=%s&start_time=%s"
	// The CandlesLimitPeriodEndpointFormat is used as GET to fetch candlesticks for a certain market with a specified limit within a certain period
	CandlesLimitPeriodEndpointFormat = "/markets/%s/%s/candles?resolution=%s&limit=%s&start_time=%s&end_time=%s"

	/*
		Futures Endpoints
	*/

	// The FuturesEndpoint is used as GET to fetch all of the available futures
	FuturesEndpoint = "/futures"
	// The ExpiredFuturesEndpoint is used as GET to fetch all of the expired futures
	ExpiredFuturesEndpoint = "/expired_futures"
	// The FuturesFundingRate is used as GET to fetch the funding rate of all futures
	FuturesFundingRate = "/funding_rates"

	/*
		Futures Endpoint Formats
	*/

	// The FuturesEndpointFormat
	FuturesEndpointFormat = "/futures/%s"
	// The FuturesStatsFormat
	FuturesStatsFormat = "/futures/%s/stats"
	// The IndexWeightsFormat
	IndexWeightsFormat = "/indexes/%s/weights"
	// The IndexCandlesEndpointFormat
	IndexCandlesEndpointFormat = "/indexes/%s/candles?resolution=%d"
	// The IndexCandlesLimitEndpointFormat
	IndexCandlesLimitEndpointFormat = "/indexes/%s/candles?resolution=%d&limit=%d"
	// The IndexCandlesLimitStartEndEndpointFormat
	IndexCandlesLimitStartEndEndpointFormat = "/indexes/%s/candles?resolution=%d&limit=%d&start_time=%d&end_time=%d"
	// The FuturesFutureFundingRate is used as GET to fetch the funding rate of a certain future
	FuturesFutureFundingRate = "/funding_rates?future=%s"
	// The FuturesFutureStartEndFundingRate is used as GET to fetch the funding rate of a certain future between a period of time
	FuturesFutureStartEndFundingRate = "/funding_rates?future=%s&start=%d&end=%d"

	/*
		Orders Endpoints
	*/

	// The OrderEndpoint is used as GET to fetch all open orders and DELETE to cancel all orders
	OrderEndpoint = "/orders"
	// The TriggerOrder
	TriggerOrder = "/conditional_orders"

	/*
		Orders Endpoint Formats
	*/

	// The OrderStatusEndpointFormat is used as GET to fetch the order status and DELETE to cancel the order
	OrderStatusEndpointFormat = "/orders/%s"
	// The OrderStatusByClientIDEndpointFormat is used as GET to fetch the orders status' and DELETE to cancel the orders
	OrderStatusByClientIDEndpointFormat = "/orders/by_client_id/%s"
	// The ModifyOrderEndpointFormat
	ModifyOrderEndpointFormat = "/orders/%s/modify"
	// The ModifyOrderByClientIDEndpointFormat
	ModifyOrderByClientIDEndpointFormat = "/orders/by_client_id/%s/modify"
	// The OpenOrdersEndpointFormat
	OpenOrdersEndpointFormat = "/orders?market=%s"
	// The OrdersHistoryEndpointFormat
	OrdersHistoryEndpointFormat = "/orders/history?market=%s"
	// The TriggerOrderCancelEndpointFormat
	TriggerOrderCancelEndpointFormat = "/conditional_orders/%s"
	// The OpenTriggerOrdersEndpointFormat
	OpenTriggerOrdersEndpointFormat = "/conditional_orders?market={market}"
	// The TriggerOrdersTriggersEndpointFormat
	TriggerOrdersTriggersEndpointFormat = "/conditional_orders/%s/triggers"
	// The TriggerOrderHistoryEndpointFormat
	TriggerOrderHistoryEndpointFormat = "/conditional_orders/history?market=%s"
	// The ModifyTriggerOrderEndpointFormat
	ModifyTriggerOrderEndpointFormat = "/conditional_orders/%s/modify"

	/*
		Convert Endpoints
	*/

	// The ConversionRequestQuoteEndpoint
	ConversionRequestQuoteEndpoint = "/otc/quotes"

	/*
		Convert Endpoint Formats
	*/

	// The ConversionQuoteStatusEndpointFormat
	ConversionQuoteStatusEndpointFormat = "/otc/quotes/%s"
	// The ConversionAcceptQuoteEndpointFormat
	ConversionAcceptQuoteEndpointFormat = "/otc/quotes/%s/accept"

	/*
		Spot Margin Endpoints
	*/

	// The MarginHistoryEndpoint
	MarginHistoryEndpoint = "/spot_margin/history"
	// The BorrowRatesEndpoint
	BorrowRatesEndpoint = "/spot_margin/borrow_rates"
	// The LendingRatesEndpoint
	LendingRatesEndpoint = "/spot_margin/lending_rates"
	// The BorrowSummaryEndpoint
	BorrowSummaryEndpoint = "/spot_margin/borrow_summary"
	// The MyBorrowHistoryEndpoint
	MyBorrowHistoryEndpoint = "/spot_margin/borrow_history"
	// The MyLendingHistoryEndpoint
	MyLendingHistoryEndpoint = "/spot_margin/lending_history"
	// The LendingOffersEndpoint is used as GET to fetch lending offers and as POST to create a new offer
	LendingOffersEndpoint = "/spot_margin/offers"
	// The LendingInfoEndpoint
	LendingInfoEndpoint = "/spot_margin/lending_info"

	/*
		Spot Margin Endpoint Formats
	*/

	// The MarketInfoEndpointFormat
	MarketInfoEndpointFormat = "/spot_margin/market_info?market=%s"

	/*
		Fills Endpoint Formats
	*/
	FillsEndpointFormat = "/fills?market=%s"

	/*
		Funding Payments Endpoints
	*/

	// The FundingPaymentsEndpoint
	FundingPaymentsEndpoint = "/funding_payments"

	/*
		Funding Payments Endpoint Formats
	*/

	// The FundingPaymentsFutureEndpoint where it only fetches for a certain future
	FundingPaymentsFutureEndpoint = "/funding_payments?future=%s"
	// The FundingPaymentsFutureStartEndEndpoint where it only fetches for a certain future between a period of time
	FundingPaymentsFutureStartEndEndpoint = "/funding_payments?future=%s&start_time=%d&end_time=%d"

	/*
		Leveraged Tokens Endpoints
	*/

	// The LeveragedTokensEndpoint
	LeveragedTokensEndpoint = "/lt/tokens"
	// The LeveragedTokenBalancesEndpoint
	LeveragedTokenBalancesEndpoint = "/lt/balances"
	// The LeveragedTokenCreationRequestsEndpoint
	LeveragedTokenCreationRequestsEndpoint = "/lt/creations"
	// The LeveragedTokenRedemptionRequestsEndpoint
	LeveragedTokenRedemptionRequestsEndpoint = "/lt/redemptions"

	/*
		Leveraged Tokens Endpoint Formats
	*/

	// The LeveragedTokenEndpointFormat
	LeveragedTokenEndpointFormat = "/lt/%s"
	// The LeveragedTokenCreationRequestEndpointFormat
	LeveragedTokenCreationRequestEndpointFormat = "/lt/%s/create"
	// The LeveragedTokenRedemptionRequestEndpointFormat
	LeveragedTokenRedemptionRequestEndpointFormat = "/lt/%s/redeem"

	/*
		Options Endpoints
	*/

	// The OptionsAccountInfoEndpoint is used as GET to fetch the quote requests and as POST to submit a new request
	OptionsAccountInfoEndpoint = "/options/account_info"
	// The OptionsTradesEndpoint
	OptionsTradesEndpoint = "/options/trades"
	// The OptionsFillsEndpoint
	OptionsFillsEndpoint = "/options/fills"
	// The OptionsPositionsEndpoint
	OptionsPositionsEndpoint = "/options/positions"
	// The OptionsVolumeEndpoint
	OptionsVolumeEndpoint = "/stats/24h_options_volume"
	// The OptionQuoteRequestsEndpoint
	OptionQuoteRequestsEndpoint = "/options/requests"
	// The MyOptionQuotesEndpoint
	MyOptionQuotesEndpoint = "/options/my_quotes"
	// The MyOptionQuoteRequestsEndpoint
	MyOptionQuoteRequestsEndpoint = "/options/my_requests"

	/*
		Options Endpoint Formats
	*/

	// The OptionQuoteRequestCancellationEndpointFormat
	OptionQuoteRequestCancellationEndpointFormat = "/options/requests/%s"
	// The OptionQuoteRequestQuotesEndpointFormat is used as GET to fetch quotes for the quote request and as POST to submit a new quote
	OptionQuoteRequestQuotesEndpointFormat = "/options/requests/%s/quotes"
	// The OptionQuoteCancelEndpointFormat
	OptionQuoteCancelEndpointFormat = "/options/quotes/%s"
	// The OptionQuoteAcceptEndpointFormat
	OptionQuoteAcceptEndpointFormat = "/options/quotes/%s/accept"
)
