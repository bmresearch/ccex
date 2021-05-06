package builder

import (
	"fmt"
	. "github.com/murlokito/ccex"

	"github.com/murlokito/ccex/config"
)

// NewExchange returns a configured client with the passed config.
func NewExchange(exchange Exchange, config *config.Configuration) (*ExchangeClient, error) {
	switch exchange {
	case Binance:
		return nil, ErrExchangeNotImplemented
	case BinanceUS:
		return nil, ErrExchangeNotImplemented
	case BitMEX:
		return nil, ErrExchangeNotImplemented
	case Bybit:
		return nil, ErrExchangeNotImplemented
	case Deribit:
		return nil, ErrExchangeNotImplemented
	case FTX:
		return NewFTXClient(config)
	case FTXUS:
		return nil, ErrExchangeNotImplemented
	default:
		return nil, fmt.Errorf("new clients error [%v]", Exchanges[exchange])
	}
}
