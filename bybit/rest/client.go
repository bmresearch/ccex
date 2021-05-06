package rest

import (
	"golang.org/x/time/rate"

	"github.com/murlokito/ccex/auth"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/internal/rest"
)

// Client
type Client struct {
	auth    *auth.Authentication
	client  *rest.Client
	limiter *rate.Limiter
}

// GetFutures
func (c *Client) GetFutures() error {
	return nil
}

// GetMarkets
func (c *Client) GetMarkets() error {
	return nil
}

// GetCandles
func (c *Client) GetCandles(symbol string, resolution int) error {
	return nil
}

// GetCandlesLimit
func (c *Client) GetCandlesLimit(symbol string, resolution int, limit int) error {
	return nil
}

// NewClient returns a new Rest Client for bybit
func NewClient(cfg config.Configuration) (*Client, error) {
	rc, err := rest.New(cfg, ApiUrl)
	if err != nil {
		return &Client{}, err
	}

	client := Client{
		client:  rc,
		limiter: rate.NewLimiter(30, 5),
	}

	return &client, nil
}
