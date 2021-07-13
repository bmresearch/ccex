package ftx

import (
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/ftx/rest"
	"github.com/murlokito/ccex/ftx/websocket"
)

// NewClientWith returns a new configured client for FTX, to be used with the agnostic builder.
func NewClientWith(config *config.Configuration, tickerHandler exchange.TickerHandler,
	tradesHandler exchange.TradesHandler, orderBookSnapshotHandler exchange.OrderBookSnapshotHandler,
	orderBookDeltaHandler exchange.OrderBookDeltaHandler) (*exchange.Client, error) {

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
	wsClient, err := websocket.NewClientWith(config, tickerHandler, tradesHandler, orderBookSnapshotHandler, orderBookDeltaHandler)
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

// NewClient returns a new configured client for FTX, to be used with the agnostic builder.
func NewClient(config *config.Configuration, messageHandler exchange.MessageHandler) (*exchange.Client, error) {

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
	wsClient, err := websocket.NewClient(config, messageHandler)
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

