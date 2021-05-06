package builder

import (
	"github.com/murlokito/ccex"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/ftx"
	"github.com/murlokito/ccex/ftx/rest"
	"github.com/murlokito/ccex/ftx/websocket"
)

// NewFTXClient returns a new configured client for FTX, to be used with the agnostic builder.
func NewFTXClient(config *config.Configuration) (*ccex.ExchangeClient, error){

	// Initialize the base http client that takes care of authentication and rate limiting
	client, err := rest.NewClient(config)
	if err != nil {
		return nil, err
	}

	// Initialize the clients for the specific API segments
	accountClient, err := ftx.NewAccountClient(client)
	if err != nil {
		return nil, err
	}

	walletClient, err := ftx.NewWalletClient(client)
	if err != nil {
		return nil, err
	}

	ordersClient, err := ftx.NewOrdersClient(client)
	if err != nil {
		return nil, err
	}

	// Initialize the websocket client
	wsClient, err := websocket.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ccex.ExchangeClient{
		Account:   accountClient,
		Wallet:    walletClient,
		Orders:    ordersClient,
		Websocket: wsClient,
	}, nil
}
