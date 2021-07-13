package websocket

import (
	"encoding/json"
	"fmt"
	"time"

	ws "github.com/murlokito/ccex/ftx/models/websocket"

	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/internal/logger"
	"github.com/murlokito/ccex/log"

	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/internal/websocket"
)

// Client represents the websocket client for FTX
type Client struct {
	// OnOrderBookDelta holds the handler for order book messages.
	OnOrderBookDelta exchange.OrderBookDeltaHandler
	// OnOrderBookSnapshot holds the handler for order book messages.
	OnOrderBookSnapshot exchange.OrderBookSnapshotHandler
	// OnTicker holds the handler for ticker messages.
	OnTicker exchange.TickerHandler
	// OnTrade holds the handler for trade messages.
	OnTrades exchange.TradesHandler

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

// OnMessage is called by the underlying websocket client whenever it reads a message, similar to event-based actions.
func (c Client) OnMessage(message []byte) error {
	var v map[string]interface{}

	err := json.Unmarshal(message, &v)
	if err != nil {
		logger.Error(err.Error())
	}

	var (
		msgChannel, msgMarket, msgType interface{}
	)

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

	msgChannel, ok = v["channel"]
	if !ok {
		c.logger.Error("Could not get message channel")
		return fmt.Errorf("could not get message channel")
	}
	channel := fmt.Sprintf("%v", msgChannel)

	msgMarket, ok = v["market"]
	if !ok {
		c.logger.Error("Could not get message market")
		return fmt.Errorf("could not get message market")
	}
	market := fmt.Sprintf("%v", msgMarket)

	switch msgType {
	case Subscribed:
		c.subscriptions[channel] = append(c.subscriptions[channel], market)
		c.logger.Infof("Successfully %v to channel {%v} for market {%v}", msgType, channel, market)
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
		c.logger.Infof("Successfully %v to channel {%v} for market {%v}", msgType, channel, market)
		break
	}

	switch channel {
	case Trades:
		if c.OnTrades != nil {
			var trades ws.TradeMessage
			err = json.Unmarshal(message, &trades)
			if err != nil {
				return err
			}
			c.OnTrades(trades.Market, trades.Data.Standard())
		}
		break
	case Orderbook:
		var orderBook ws.OrderBookMessage
		err = json.Unmarshal(message, &orderBook)
		//fmt.Printf("action: %v \nbids: %v\nasks: %v\n", orderBook.Data.Action, orderBook.Data.Bids, orderBook.Data.Asks)
		if err != nil {
			return err
		}
		switch orderBook.Data.Action {
		case Snapshot:
			if c.OnOrderBookSnapshot != nil {
				c.OnOrderBookSnapshot(orderBook.Market, orderBook.Snapshot())
			}
			break
		case Delta:
			if c.OnOrderBookDelta != nil {
				c.OnOrderBookDelta(orderBook.Market, orderBook.Delta())
			}
			break
		}
		break
	case Ticker:
		if c.OnTicker != nil {
			var ticker ws.TickerMessage
			err = json.Unmarshal(message, &ticker)
			if err != nil {
				return err
			}
			c.OnTicker(ticker.Market, ticker.Data.Standard())
		}
		break
	}
	return nil
}

// Authenticate is called by the underling websocket client whenever it connects/reconnects in order to access private channels.
func (c Client) Authenticate() ([]byte, error) {
	ts, signature := c.AuthenticationMessage()

	data := ws.LoginMessage{
		BaseOperation: ws.BaseOperation{
			Op: "login",
		},
		AuthenticationMessage: ws.AuthenticationMessage{
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

	data := ws.SubscribeMessage{
		BaseOperation: ws.BaseOperation{
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

	return nil
}

// NewClientWith returns a configured websocket client for FTX
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
		OnOrderBookDelta:    orderBookDeltaHandler,
		OnOrderBookSnapshot: orderBookSnapshotHandler,
	}

	if config != nil {
		if config.GetAuth() != nil {
			websocketBase.OnConnected = client.Authenticate
		}
	}

	websocketBase.OnMessage = client.OnMessage

	websocketBase.SetKeepAliveTimeout(15 * time.Second)

	return client, nil
}

// NewClient returns a configured websocket client for FTX
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

	websocketBase.SetKeepAliveTimeout(15 * time.Second)

	return client, nil
}
