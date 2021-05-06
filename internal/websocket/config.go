package websocket

import "github.com/murlokito/ccex/auth"

type Config interface {
	Auth() *auth.Authentication
	BaseUrl() string
}
