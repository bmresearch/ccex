package exchange

type (
	// Exchange represents the exchange name.
	Exchange int

	/*
		ExchangeClient exposes a unified API to interact with the exchange.
		In certain cases some methods will not be implemented and thus will
		return an error accordingly.
	*/
	Client struct {
		// Account interface which interacts with account generic endpoints
		Account Account

		// Account interface which interacts with sub-account specific endpoints
		SubAccount SubAccount

		// Wallet interface which interacts with wallet-specific endpoints
		Wallet Wallet

		// Orders interface which interacts with order-specific endpoints
		Orders Orders

		// Markets interface which interacts with market generic endpoints
		Markets Markets

		// Spot interface which interacts with spot-specific endpoints
		Spot Spot

		// Futures interface which interacts with future-specific endpoints
		Futures Futures

		// Options interface which interacts with options-specific endpoints
		Options Options

		// Streaming interface which interacts with the websocket
		Streaming Websocket
	}
)
