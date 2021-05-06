package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/internal/logger"
	"github.com/murlokito/ccex/log"

	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/ftx/websocket/models"
	"github.com/murlokito/ccex/internal/websocket"
)

// Client represents the websocket client for FTX
type Client struct {
	// OnMarkets holds the handler for markets messages.
	OnMarkets exchange.OnMarketsHandler
	// OnOrderBook holds the handler for order book messages.
	OnOrderBook exchange.OnOrderBookHandler
	// OnTicker holds the handler for ticker messages.
	OnTicker exchange.OnTickerHandler
	// OnTrade holds the handler for trade messages.
	OnTrade exchange.OnTradeHandler

	// config holds the config used to establish the connection.
	config *config.Configuration

	// ws holds the underlying websocket connection.
	ws *websocket.Client

	// logger holds a logger, the user can inject a logger as long as it implements the interface we specify.
	logger log.Logger

	// subscriptions holds all subscriptions across all channels and markets.
	subscriptions map[string][]string
}

func (c *Client) Subscriptions() map[string][]string {
	return c.subscriptions
}

// Connect performs the connection
func (c *Client) Connect() {
	c.ws.Dial()
}

// Reconnect attempts reconnection
func (c *Client) Reconnect() error {
	return c.ws.CloseAndReconnect()
}

// Connected returns a boolean representing the connection state
func (c *Client) Connected() bool {
	return c.ws.Connected()
}

// OnMarketHandler sets the handler for market messages
func (c *Client) OnMarketHandler(handler exchange.OnMarketsHandler) {
	c.OnMarkets = handler
}

// OnOrderBookHandler sets the handler for orderbook messages
func (c *Client) OnOrderBookHandler(handler exchange.OnOrderBookHandler) {
	c.OnOrderBook = handler
}

// OnTradesHandler sets the handler for trade messages
func (c *Client) OnTradesHandler(handler exchange.OnTradeHandler) {
	c.OnTrade = handler
}

// OnTickerHandler sets the handler for ticker messages
func (c *Client) OnTickerHandler(handler exchange.OnTickerHandler) {
	c.OnTicker = handler
}

// OnMessage is called by the underlying websocket client whenever it reads a message, similar to event-based actions.
func (c Client) OnMessage(message []byte) error {
	var v map[string]interface{}

	err := json.Unmarshal(message, &v)
	if err != nil {
		logger.Error(err.Error())
	}

	msgType, ok := v["type"]
	if !ok {
		return fmt.Errorf("could not get message type")
	}

	if msgType == "error" {
		code, ok := v["code"]
		if !ok {
			return fmt.Errorf("could not get message code")
		}

		msg, ok := v["msg"]
		if !ok {
			return fmt.Errorf("could not get message")
		}

		return fmt.Errorf("code: %v type: %v msg: %v", code, msgType, msg)
	}
	var (
		channel, market interface{}
	)

	channel, ok = v["channel"]
	if !ok {
		c.logger.Error("Could not get message channel")
		return fmt.Errorf("could not get message channel")
	}

	market, ok = v["market"]
	if !ok {
		c.logger.Error("Could not get message market")
		return fmt.Errorf("could not get message market")
	}

	if msgType == "subscribed" || msgType == "unsubscribed" {
		c.logger.Infof("Successfully %v to channel {%v} for market {%v}", msgType, channel, market)
		return nil
	}

	switch channel {
	case Markets:
		if c.OnMarkets != nil {
			var markets models.MarketMessage
			err = json.Unmarshal(message, &markets)
			if err != nil {
				return err
			}
			c.OnMarkets(markets)
		}
		break
	case Trades:
		if c.OnTrade != nil {
			var trades models.TradeMessage
			err = json.Unmarshal(message, &trades)
			if err != nil {
				return err
			}
			c.OnTrade(trades)
		}
		break
	case Orderbook:
		if c.OnOrderBook != nil {
			var orderbook models.OrderBookMessage
			err = json.Unmarshal(message, &orderbook)
			if err != nil {
				return err
			}
			c.OnOrderBook(orderbook)
		}
		break
	case Ticker:
		if c.OnTicker != nil {
			var ticker models.TickerMessage
			err = json.Unmarshal(message, &ticker)
			if err != nil {
				return err
			}
			c.OnTicker(ticker)
		}
		break
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
	err = c.ws.WriteMessage(1, message)
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

	err = c.ws.WriteMessage(1, message)
	if err != nil {
		return err
	}

	c.subscriptions[channel] = append(c.subscriptions[channel], market)

	return nil
}

// NewClient returns a configured websocket client for FTX
func NewClient(config *config.Configuration, marketsHandler exchange.OnMarketsHandler,
	tickerHandler exchange.OnTickerHandler, tradesHandler exchange.OnTradeHandler, orderbookHandler exchange.OnOrderBookHandler) (*Client, error) {
	clientLogger := logger.NewLogger()
	ws, err := websocket.New(Url, clientLogger)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config:        config,
		ws:            ws,
		logger:        clientLogger,
		subscriptions: map[string][]string{},
		OnTicker:      tickerHandler,
		OnTrade:       tradesHandler,
		OnOrderBook:   orderbookHandler,
		OnMarkets:     marketsHandler,
	}

	if config != nil {
		if config.GetAuth() != nil {
			ws.OnConnected = client.Authenticate
		}
	}

	ws.OnMessage = client.OnMessage

	ws.SetKeepAliveTimeout(15 * time.Second)

	return client, nil
}
