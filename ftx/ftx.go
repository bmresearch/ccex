package ftx

import (
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/ftx/rest"
	"github.com/murlokito/ccex/ftx/websocket"
)

// NewFTXClient returns a new configured client for FTX, to be used with the agnostic builder.
func NewFTXClient(config *config.Configuration, marketsHandler exchange.OnMarketsHandler,
	tickerHandler exchange.OnTickerHandler, tradesHandler exchange.OnTradeHandler,
	orderBookHandler exchange.OnOrderBookHandler) (*exchange.Client, error) {

	// Initialize the base http client that takes care of authentication and rate limiting
	client, err := rest.NewClient(config)
	if err != nil {
		return nil, err
	}

	// Initialize the clients for the specific API segments
	accountClient, err := NewAccountClient(client)
	if err != nil {
		return nil, err
	}

	subAccountClient, err := NewSubAccountClient(client)
	if err != nil {
		return nil, err
	}

	walletClient, err := NewWalletClient(client)
	if err != nil {
		return nil, err
	}

	ordersClient, err := NewOrdersClient(client)
	if err != nil {
		return nil, err
	}

	marketsClient, err := NewMarketsClient(client)
	if err != nil {
		return nil, err
	}

	spotClient, err := NewSpotClient(client)
	if err != nil {
		return nil, err
	}

	futuresClient, err := NewFuturesClient(client)
	if err != nil {
		return nil, err
	}

	optionsClient, err := NewOptionsClient(client)
	if err != nil {
		return nil, err
	}

	// Initialize the websocket client
	wsClient, err := websocket.NewClient(config, marketsHandler, tickerHandler, tradesHandler, orderBookHandler)
	if err != nil {
		return nil, err
	}

	return &exchange.Client{
		Account:    accountClient,
		SubAccount: subAccountClient,
		Wallet:     walletClient,
		Orders:     ordersClient,
		Markets:    marketsClient,
		Spot:       spotClient,
		Futures:    futuresClient,
		Options:    optionsClient,
		Streaming:  wsClient,
	}, nil
}
