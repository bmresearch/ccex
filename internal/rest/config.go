package rest

import "github.com/murlokito/ccex/auth"

type Config interface {
	Auth() *auth.Authentication
	SubAccount() string
}
