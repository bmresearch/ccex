package websocket

type (
	// BaseOperation is the base of the message, all common attributes.
	BaseOperation struct {
		Op string `json:"op"`
	}

	// AuthenticationMessage holds the authentication-relevant information.
	AuthenticationMessage struct {
		Key       string `json:"key"`
		Signature string `json:"sign"`
		Timestamp string `json:"time"`
	}

	// LoginMessage is used to authenticate with the websocket server.
	LoginMessage struct {
		BaseOperation
		AuthenticationMessage `json:"args"`
	}

	// SubscribeMessage is used to request a subscription to a channel.
	SubscribeMessage struct {
		BaseOperation
		Channel string `json:"channel"`
		Market  string `json:"market"`
	}
)
