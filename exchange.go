package ccex

type (
	// Exchange represents the exchange name.
	Exchange int

	/*
		ExchangeClient exposes a unified API to interact with the exchange.
		In certain cases some methods will not be implemented and thus will
		return an error accordingly.
	*/
	ExchangeClient struct {
		Account Account

		Wallet Wallet

		Orders Orders

		Websocket Websocket
	}

)
