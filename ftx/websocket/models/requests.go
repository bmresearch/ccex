package models

import (
	"github.com/murlokito/ccex"
)

// BaseOperation is the base of the message, all common attributes.
type BaseOperation struct {
	Op string `json:"op"`
}

// AuthenticationMessage holds the authentication-relevant information.
type AuthenticationMessage struct {
	Key       string `json:"key"`
	Signature string `json:"sign"`
	Timestamp string `json:"time"`
}

// LoginMessage is used to authenticate with the websocket server.
type LoginMessage struct {
	BaseOperation
	AuthenticationMessage `json:"args"`
}

// SubscribeMessage is used to request a subscription to a channel.
type SubscribeMessage struct {
	BaseOperation
	Channel ccex.Channel `json:"channel"`
	Market  ccex.Market  `json:"market"`
}
