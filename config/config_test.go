package config

import (
	"testing"

	"github.com/murlokito/ccex/auth"
)

func TestConfig(t *testing.T) {
	config := Configuration{
		Auth: &auth.Authentication{
			Key:    "testkey",
			Secret: "testsecret",
		},
		SubAccount: "none",
	}

	t.Run("have key and secret", func(t *testing.T) {

		if config.GetAuth() == nil {
			t.Error("auth is nil")
		}
	})
}

func TestSetConfig(t *testing.T) {
	config := Configuration{
		Auth: nil,
		SubAccount: "none",
	}
	authInfo := &auth.Authentication{
		Key:    "testkey",
		Secret: "testsecret",
	}
	t.Run("set key and secret", func(t *testing.T) {

		config.SetAuth(authInfo)
		if config.GetAuth() == nil {
			t.Error("failed to set auth")
		}

	})
}