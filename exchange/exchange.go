package exchange

type (
	// Exchange represents the exchange name.
	Exchange int

	/*
		ExchangeClient exposes a unified API to interact with the exchange.
		In certain cases some methods will not be implemented and thus will
		return an error accordingly.
	*/
	ExchangeClient struct {
		// Account interface which interacts with account generic endpoints
		Account Account

		// Wallet interface which interacts with wallet-specific endpoints
		Wallet Wallet

		// Conversion interface which interacts with conversion-specific endpoints
		Conversion Conversion

		// Orders interface which interacts with order-specific endpoints
		Orders Orders

		// Markets interface which interacts with market generic endpoints
		Markets Markets

		// Fills interface which interacts with order fill specific endpoints
		Fills Fills

		// Funding interface which interacts with funding specific endpoints
		Funding Funding

		// Spot interface which interacts with spot-specific endpoints
		Spot Spot

		// Futures interface which interacts with future-specific endpoints
		Futures Futures

		// Margin interface which interacts with spot margin-specific endpoints
		Margin Margin

		// Options interface which interacts with options-specific endpoints
		Options Options

		// Websocket interface which interacts with the websocket
		Websocket Websocket
	}
)
