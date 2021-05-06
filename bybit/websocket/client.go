package websocket

import (
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/internal/websocket"
)

// Client
type Client struct {
	ws *websocket.Client
}

// NewClient returns a new websocket client for bybit
func NewClient(config config.Configuration) (*Client, error) {
	ws, err := websocket.New(config, WebsocketUrl)
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		ws: ws,
	}

	return client, nil
}
