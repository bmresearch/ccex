package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/murlokito/ccex"
	"github.com/murlokito/ccex/internal/logger"
	"github.com/murlokito/ccex/log"

	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/ftx/websocket/models"
	"github.com/murlokito/ccex/internal/websocket"
)

// Client represents the websocket client for FTX
type Client struct {
	// OnMarkets holds the handler for markets messages.
	OnMarkets   ccex.OnMarketsHandler
	// OnOrderBook holds the handler for order book messages.
	OnOrderBook ccex.OnOrderBookHandler
	// OnTicker holds the handler for ticker messages.
	OnTicker    ccex.OnTickerHandler
	// OnTrade holds the handler for trade messages.
	OnTrade     ccex.OnTradeHandler

	// config holds the config used to establish the connection.
	config *config.Configuration

	// ws holds the underlying websocket connection.
	ws     *websocket.Client

	// logger holds a logger, the user can inject a logger as long as it implements the interface we specify.
	logger log.Logger

	// subscriptions holds all subscriptions across all channels and markets.
	subscriptions map[string]string

	// subscriptions holds all subscriptions with personalized handlers.
	subscriptionHandlers []ccex.MessageDispatcher
}

func (c *Client) MarketHandler(handler ccex.OnMarketsHandler) {
	c.OnMarkets = handler
}

func (c *Client) OrderBookHandler(handler ccex.OnOrderBookHandler) {
	c.OnOrderBook = handler
}

func (c *Client) TradeHandler(handler ccex.OnTradeHandler) {
	c.OnTrade = handler
}

func (c *Client) TickerHandler(handler ccex.OnTickerHandler) {
	c.OnTicker = handler
}

// OnMessage is called by the underlying websocket client whenever it reads a message, similar to event-based actions.
func (c Client) OnMessage(message []byte) {
	var v map[string]string

	err := json.Unmarshal(message, v)
	if err != nil {
		logger.Error(fmt.Errorf(""))
	}

	channel, ok := v["channel"]
	if !ok {
		c.logger.Error("Could not get message channel")
		return
	}

	market, ok := v["market"]
	if !ok {
		c.logger.Error("Could not get message channel")
		return
	}

	if handler := c.GetHandlerFor(channel, market); handler != nil {
		handler(v)
		return
	}

	switch channel {
		case Markets:
			if c.OnMarkets != nil {
				var markets models.MarketMessage
				err := json.Unmarshal(message, markets)
				if err != nil {
					logger.Error(fmt.Errorf(""))
				}
				c.OnMarkets(markets)
			}
			break
		case Trades:
			if c.OnTrade != nil {
				var trades models.TradeMessage
				err := json.Unmarshal(message, trades)
				if err != nil {
					logger.Error(fmt.Errorf(""))
				}
				c.OnTrade(trades)
			}
			break
		case Orderbook:
			if c.OnOrderBook != nil {
				var orderbook models.OrderBookMessage
				err := json.Unmarshal(message, orderbook)
				if err != nil {
					logger.Error(fmt.Errorf(""))
				}
				c.OnOrderBook(orderbook)
			}
			break
		case Ticker:
			if c.OnTicker != nil {
				var ticker models.TickerMessage
				err := json.Unmarshal(message, ticker)
				if err != nil {
					logger.Error(fmt.Errorf(""))
				}
				c.OnTicker(ticker)
			}
			break
	}

}

// GetHandlerFor fetches the personalized handler for a channel and market.
func (c Client) GetHandlerFor(channel string, market string) ccex.OnMessageHandler {
	for _, dispatcher := range c.subscriptionHandlers {
		if dispatcher.Channel == channel && market == dispatcher.Market {
			return dispatcher.Handler
		}
	}
	return nil
}

// Authenticate is called by the underling websocket client whenever it connects/reconnects in order to access private channels.
func (c Client) Authenticate() ([]byte, error) {
	ts, signature := c.AuthenticationMessage()

	data := models.LoginMessage{
		BaseOperation: models.BaseOperation{
			Op: "login",
		},
		AuthenticationMessage: models.AuthenticationMessage{
			Key:       c.config.Auth.Secret,
			Signature: signature,
			Timestamp: ts,
		},
	}

	message, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = c.ws.WriteMessage(2, message)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Subscribe subscribes to a websocket channel.
func (c Client) Subscribe(channel string, market string) error {

	data := models.SubscribeMessage{
		BaseOperation: models.BaseOperation{
			Op: "subscribe",
		},
		Channel: channel,
		Market:  market,
	}

	message, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = c.ws.WriteMessage(2, message)
	if err != nil {
		return err
	}

	return nil
}

// SubscribeWithHandler subscribes to a websocket channel.
func (c Client) SubscribeWithHandler(channel string, market string, handler ccex.OnMessageHandler) error {

	data := models.SubscribeMessage{
		BaseOperation: models.BaseOperation{
			Op: "subscribe",
		},
		Channel: channel,
		Market:  market,
	}

	message, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = c.ws.WriteMessage(2, message)
	if err != nil {
		return err
	}

	msgDispatcher := ccex.MessageDispatcher{
		Channel: channel,
		Market:  market,
		Handler: handler,
	}
	c.subscriptionHandlers = append(c.subscriptionHandlers, msgDispatcher)

	return nil
}

// NewClient returns a configured websocket client for FTX
func NewClient(config *config.Configuration) (*Client, error) {
	ws, err := websocket.New(Url)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config:               config,
		ws:                   ws,
		subscriptionHandlers: []ccex.MessageDispatcher{},
	}

	if config.GetAuth() != nil {
		ws.OnConnected = client.Authenticate
	}

	ws.OnMessage = client.OnMessage

	return client, nil
}
