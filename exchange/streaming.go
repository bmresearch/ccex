package exchange

import (
	"github.com/murlokito/ccex/models/ws"
)

type (

	/*
		OrderBookDeltaHandler is called whenever the websocket client receives an order book delta message.

		This message specifies points in the order book which must be updated, there may be insertions, updates and deletions.
	 */
	OrderBookDeltaHandler func(symbol string, data ws.OrderBookDelta)

	/*
		OrderBookSnapshotHandler is called whenever the websocket client receives an order book snapshot message.

		This message establishes the initial state of the order book, all subsequent messages will be dispatched by the OrderBookDeltaHandler.
	 */
	OrderBookSnapshotHandler func(symbol string, data ws.OrderBookSnapshot)

	// TickerHandler is called whenever the websocket client receives a message from the ticker/instrument channels.
	TickerHandler func(symbol string, data ws.Ticker)

	// TradesHandler is called whenever the websocket client receives a message from the trades channel.
	TradesHandler func(symbol string, data ws.Trades)

	/*
		MessageHandler is a type defined to represent a handler called when a message is read from the connection.

		This allows the user of the library to, in a way, override the internal message handler.
		When the method which allows this parameter is used, the user will need to process the messages accordingly.
	*/
	MessageHandler func(message []byte) error

	// Websocket specifies functionality to interact with the websocket API.
	Websocket interface {
		Connect()
		Connected() bool
		Reconnect() error
		Subscriptions() map[string][]string
		Subscribe(channel string, market string) error
	}
)
