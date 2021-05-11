package websocket

import (
	"encoding/json"
	"strings"
	"time"

	ws "github.com/murlokito/ccex/bybit/models/websocket"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/internal/logger"
	"github.com/murlokito/ccex/internal/websocket"
	"github.com/murlokito/ccex/log"
)

// Client represents the websocket client for Bybit
type Client struct {
	// OnOrderBookSnapshot holds the handler for order book snapshot messages.
	OnOrderBookSnapshot exchange.OrderBookSnapshotHandler
	// OnOrderBookSnapshot holds the handler for order book delta messages.
	OnOrderBookDelta exchange.OrderBookDeltaHandler
	// OnTicker holds the handler for ticker messages.
	OnTicker exchange.TickerHandler
	// OnTrades holds the handler for trade messages.
	OnTrades exchange.TradesHandler

	// config holds the config used to establish the connection.
	config *config.Configuration

	// ws holds the underlying websocket connection.
	ws *websocket.Client

	// logger holds a logger, the user can inject a logger as long as it implements the interface we specify.
	logger log.Logger

	// subscriptions holds all subscriptions across all channels and markets.
	subscriptions map[string][]string

	instruments map[string]*ws.Instrument
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

// OnMessage is called by the underlying websocket client whenever it reads a message, similar to event-based actions.
func (c Client) OnMessage(message []byte) error {
	var (
		v ws.SubscriptionResponse
	)

	err := json.Unmarshal(message, &v)
	if err != nil {
		// in case an error happens just assume the message is due to a subscription and try to process it
		return c.ProcessMessage(message)
	}

	split := strings.Split(v.Request.Args[0], ".")
	channel := split[0]
	market := split[1]

	switch v.Request.Op {
	case Subscribed:
		c.subscriptions[channel] = append(c.subscriptions[channel], market)
		c.logger.Infof("Successfully %v to channel {%v} for market {%v}", v.Request.Op, channel, market)
		break
	case Unsubscribed:
		var subs []string
		for _, sub := range c.subscriptions[channel] {
			if sub == market {
				continue
			}
			subs = append(subs, sub)
		}
		c.subscriptions[channel] = subs
		c.logger.Infof("Successfully %v to channel {%v} for market {%v}", v.Request.Op, channel, market)
		break
	}

	return nil
}

// ProcessMessage attempts to process a non-subscription response message
func (c Client) ProcessMessage(message []byte) error {
	var (
		msg ws.BaseDataResponse
	)

	err := json.Unmarshal(message, &msg)
	if err != nil {
		return err
	}

	// Order book data
	if strings.HasPrefix(msg.Topic, OrderBook) {
		var symbol string

		// Order book with a depth of 25 orders per side
		if strings.HasPrefix(msg.Topic, OrderBook25) {
			symbol = strings.TrimLeft(msg.Topic, OrderBook25+".")
		}
		// Order book with a depth of 200 orders per side
		if strings.HasPrefix(msg.Topic, OrderBook200) {
			symbol = strings.TrimLeft(msg.Topic, OrderBook200+".")
		}

		switch msg.Type {
		case Snapshot:
			if c.OnOrderBookSnapshot != nil{
				var data *ws.OrderBookSnapshot
				err = json.Unmarshal(msg.Data.([]byte), &data)
				if err != nil {
					return err
				}
				c.OnOrderBookSnapshot(symbol, data.Standard())
			}
			break
		case Delta:
			if c.OnOrderBookDelta != nil {
				var data *ws.OrderBookDelta
				err = json.Unmarshal(msg.Data.([]byte), &data)
				if err != nil {
					return err
				}
				c.OnOrderBookDelta(symbol, data.Standard())
			}
			break
		}
	}

	if strings.HasPrefix(msg.Topic, Instrument) {
		symbol := strings.TrimLeft(msg.Topic, Instrument)

		switch msg.Type {
		case Snapshot:
			if c.OnTicker != nil {
				var data *ws.InstrumentSnapshot
				err = json.Unmarshal(msg.Data.([]byte), &data)
				if err != nil {
					return err
				}
				c.instruments[symbol] = data.Instrument()
				c.OnTicker(symbol, data.Standard())
			}
			break
		case Delta:
			if c.OnTicker != nil {
				var data *ws.InstrumentDelta
				err = json.Unmarshal(msg.Data.([]byte), &data)
				if err != nil {
					return err
				}
				c.instruments[symbol].Update(data)
				c.OnTicker(symbol, c.instruments[symbol].Standard())
			}
			break
		}
	}

	if strings.HasPrefix(msg.Topic, Trades) {
		if c.OnTrades != nil {
			symbol := strings.TrimLeft(msg.Topic, Trades)
			var data *ws.Trades
			err = json.Unmarshal(msg.Data.([]byte), &data)
			if err != nil {
				return err
			}
			c.OnTrades(symbol, data.Standard())
		}
	}

	return nil
}

// Authenticate is called by the underling websocket client whenever it connects/reconnects in order to access private channels.
func (c Client) Authenticate() ([]byte, error) {
	ts, signature := c.AuthenticationMessage()

	data := ws.Message{
		BaseOperation: ws.BaseOperation{
			Op: "auth",
		},
		Args: []string{
			c.config.Auth.Key,
			ts,
			signature,
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

	data := ws.Message{
		BaseOperation: ws.BaseOperation{
			Op: "subscribe",
		},
		Args: []string{},
	}

	message, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = c.ws.WriteMessage(1, message)
	if err != nil {
		return err
	}

	return nil
}

// NewClientWith returns a new websocket client for bybit
func NewClientWith(config *config.Configuration, tickerHandler exchange.TickerHandler,
	tradesHandler exchange.TradesHandler, orderBookSnapshotHandler exchange.OrderBookSnapshotHandler,
	orderBookDeltaHandler exchange.OrderBookDeltaHandler) (*Client, error) {

	clientLogger := logger.NewLogger()

	websocketBase, err := websocket.New(Url, clientLogger)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config:              config,
		ws:                  websocketBase,
		logger:              clientLogger,
		subscriptions:       map[string][]string{},
		OnTicker:            tickerHandler,
		OnTrades:            tradesHandler,
		OnOrderBookSnapshot: orderBookSnapshotHandler,
		OnOrderBookDelta:    orderBookDeltaHandler,
	}

	if config != nil {
		if config.GetAuth() != nil {
			websocketBase.OnConnected = client.Authenticate
		}
	}

	websocketBase.OnMessage = client.OnMessage

	websocketBase.SetKeepAliveTimeout(30 * time.Second)

	return client, nil
}

// NewClient returns a new websocket client for bybit
func NewClient(config *config.Configuration, messageHandler exchange.MessageHandler) (*Client, error) {

	clientLogger := logger.NewLogger()

	websocketBase, err := websocket.New(Url, clientLogger)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config:        config,
		ws:            websocketBase,
		logger:        clientLogger,
		subscriptions: map[string][]string{},
	}

	if config != nil {
		if config.GetAuth() != nil {
			websocketBase.OnConnected = client.Authenticate
		}
	}

	websocketBase.OnMessage = websocket.OnMessageHandler(messageHandler)

	websocketBase.SetKeepAliveTimeout(30 * time.Second)

	return client, nil
}
