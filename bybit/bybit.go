package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	"github.com/murlokito/ccex/bybit/websocket"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/exchange"
)

// NewClientWith returns a new configured client for Bybit, to be used with the agnostic builder.
func NewClientWith(config *config.Configuration, tickerHandler exchange.TickerHandler,
	tradesHandler exchange.TradeHandler, orderBookHandler exchange.OrderBookHandler) (*exchange.Client, error) {

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

	futuresClient, err := NewFuturesClient(client)
	if err != nil {
		return nil, err
	}

	// Initialize the websocket client
	wsClient, err := websocket.NewClientWith(config, tickerHandler, tradesHandler, orderBookHandler)
	if err != nil {
		return nil, err
	}

	return &exchange.Client{
		Account:    accountClient,
		SubAccount: nil,
		Wallet:     walletClient,
		Orders:     ordersClient,
		Markets:    marketsClient,
		Spot:       nil,
		Futures:    futuresClient,
		Options:    nil,
		Streaming:  wsClient,
	}, nil
}

// NewClient returns a new configured client for Bybit, with a defined message handler.
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

	futuresClient, err := NewFuturesClient(client)
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
		SubAccount: nil,
		Wallet:     walletClient,
		Orders:     ordersClient,
		Markets:    marketsClient,
		Spot:       nil,
		Futures:    futuresClient,
		Options:    nil,
		Streaming:  wsClient,
	}, nil
}
