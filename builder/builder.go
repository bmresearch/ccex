package builder

import (
	"fmt"
	"github.com/murlokito/ccex"
	"github.com/murlokito/ccex/bybit"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/ftx"
)

// NewExchangeWith returns a configured exchange client with the passed config and message handlers.
func NewExchangeWith(
	exchange exchange.Exchange,
	config *config.Configuration,
	tickerHandler exchange.TickerHandler,
	tradesHandler exchange.TradeHandler,
	orderBookHandler exchange.OrderBookHandler) (*exchange.Client, error) {
	switch exchange {
	case Binance:
		return nil, ccex.ErrExchangeNotImplemented
	case BinanceUS:
		return nil, ccex.ErrExchangeNotImplemented
	case BitMEX:
		return nil, ccex.ErrExchangeNotImplemented
	case Bybit:
		return bybit.NewClientWith(config, tickerHandler, tradesHandler, orderBookHandler)
	case Deribit:
		return nil, ccex.ErrExchangeNotImplemented
	case FTX:
		return ftx.NewClientWith(config, tickerHandler, tradesHandler, orderBookHandler)
	case FTXUS:
		return nil, ccex.ErrExchangeNotImplemented
	default:
		return nil, fmt.Errorf("new clients error [%v]", Exchanges[exchange])
	}
}

/*
NewExchange returns a configured exchange client with the passed config and overriding the internal message processor.

When using this initialization method the user will need to process messages accordingly.
 */
func NewExchange(
	exchange exchange.Exchange,
	config *config.Configuration,
	messageHandler exchange.MessageHandler,) (*exchange.Client, error) {
	switch exchange {
	case Binance:
		return nil, ccex.ErrExchangeNotImplemented
	case BinanceUS:
		return nil, ccex.ErrExchangeNotImplemented
	case BitMEX:
		return nil, ccex.ErrExchangeNotImplemented
	case Bybit:
		return bybit.NewClient(config, messageHandler)
	case Deribit:
		return nil, ccex.ErrExchangeNotImplemented
	case FTX:
		return ftx.NewClient(config, messageHandler)
	case FTXUS:
		return nil, ccex.ErrExchangeNotImplemented
	default:
		return nil, fmt.Errorf("new clients error [%v]", Exchanges[exchange])
	}
}
