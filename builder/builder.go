package builder

import (
	"fmt"
	"github.com/murlokito/ccex"
	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/exchange"
	"github.com/murlokito/ccex/ftx"
)

// NewExchange returns a configured client with the passed config.
func NewExchange(exchange exchange.Exchange, config *config.Configuration) (*exchange.ExchangeClient, error) {
	switch exchange {
	case Binance:
		return nil, ccex.ErrExchangeNotImplemented
	case BinanceUS:
		return nil, ccex.ErrExchangeNotImplemented
	case BitMEX:
		return nil, ccex.ErrExchangeNotImplemented
	case Bybit:
		return nil, ccex.ErrExchangeNotImplemented
	case Deribit:
		return nil, ccex.ErrExchangeNotImplemented
	case FTX:
		return ftx.NewFTXClient(config)
	case FTXUS:
		return nil, ccex.ErrExchangeNotImplemented
	default:
		return nil, fmt.Errorf("new clients error [%v]", Exchanges[exchange])
	}
}
