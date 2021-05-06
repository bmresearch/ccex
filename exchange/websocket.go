package exchange

import "github.com/murlokito/ccex/ftx/websocket/models"

type (
	// OnMarketsHandler is called whenever the websocket client receives a market message.
	OnMarketsHandler func(data models.MarketMessage)

	// OnOrderBookHandler is called whenever the websocket client receives an order book message.
	OnOrderBookHandler func(data models.OrderBookMessage)

	// OnTickerHandler is called whenever the websocket client receives a ticker message.
	OnTickerHandler func(data models.TickerMessage)

	// OnTradeHandler is called whenever the websocket client receives a trade message.
	OnTradeHandler func(data models.TradeMessage)

	// OnMessageHandler is a type defined to represent a handler called for a certain channel and market combination.
	OnMessageHandler func(message interface{})

	// MessageDispatcher represents a subscription with a personal handler.
	MessageDispatcher struct {
		Channel string
		Market  string
		Handler OnMessageHandler
	}

	// Websocket specifies functionality to interact with the websocket API.
	Websocket interface {
		Connect()
		Connected() bool
		Reconnect() error
		Subscriptions() map[string][]string
		Subscribe(channel string, market string) error
		OnMarketHandler(handler OnMarketsHandler)
		OnOrderBookHandler(handler OnOrderBookHandler)
		OnTradesHandler(handler OnTradeHandler)
		OnTickerHandler(handler OnTickerHandler)
	}
)
