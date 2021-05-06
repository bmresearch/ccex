package builder

const (
	Binance = iota
	BinanceUS
	BitMEX
	Bybit
	Deribit
	FTX
	FTXUS
)

// Exchanges exposes a slice of strings in the same order as they are defined by the constants
var Exchanges = []string{
	"binance",
	"binance-us",
	"bitmex",
	"bybit",
	"deribit",
	"ftx",
	"ftx-us",
}
