package bybit

import (
	"github.com/murlokito/ccex/bybit/rest"
	"github.com/murlokito/ccex/bybit/websocket"
	"github.com/murlokito/ccex/config"
)

// Client represents the REST API Client for bybit
type Client struct {
	restClient *rest.Client
	wsClient   *websocket.Client
}

// GetMarkets fetches
func (c *Client) GetMarkets() error {
	return c.restClient.GetMarkets()
}

// GetCandles fetches
func (c *Client) GetCandles(symbol string, resolution int) error {
	return c.restClient.GetCandles(symbol, resolution)
}

// GetCandlesLimit fetches
func (c *Client) GetCandlesLimit(symbol string, resolution int, limit int) error {
	return c.restClient.GetCandlesLimit(symbol, resolution, limit)
}

// NewClient returns a configured client for bybit
func NewClient(cfg config.Configuration) (*Client, error) {

	rc, err := rest.NewClient(cfg)
	if err != nil {
		return &Client{}, err
	}

	wsc, err := websocket.NewClient(cfg)
	if err != nil {
		return &Client{}, err
	}

	c := &Client{
		restClient: rc,
		wsClient:   wsc,
	}

	return c, nil
}
