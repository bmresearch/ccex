package config

import (
	"fmt"
	"github.com/murlokito/ccex/auth"
)

// Configuration holds everything necessary to
type Configuration struct {
	Auth       *auth.Authentication
	SubAccount string
}

func (c Configuration) String() string {
	return fmt.Sprintf("k: %v s: %v sub-account: %v", c.Auth.GetKey(), c.Auth.GetSecret(), c.SubAccount)
}

// GetAuth retrieves the configuration's authentication
func (c *Configuration) GetAuth() *auth.Authentication {
	return c.Auth
}

// SetAuth sets the configuration's authentication
func (c *Configuration) SetAuth(auth *auth.Authentication) {
	c.Auth = auth
}

// GetSubAccount retrieves the configuration's sub account in use
func (c *Configuration) GetSubAccount() string {
	return c.SubAccount
}

// SetAuth sets the configuration's sub account in use
func (c *Configuration) SetSubAccount(account string) {
	c.SubAccount = account
}
