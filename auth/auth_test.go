package auth

import "testing"

func TestAuth(t *testing.T) {
	auth := Authentication{
		Key:    "testkey",
		Secret: "testsecret",
	}

	t.Run("have key and secret", func(t *testing.T) {

		if auth.GetKey() != "testkey" {
			t.Error("key is different")
		}

		if auth.GetSecret() != "testsecret" {
			t.Error("secret is different")
		}
	})
}

func TestSetAuth(t *testing.T) {
	auth := Authentication{
		Key:    "",
		Secret: "",
	}

	t.Run("set key and secret", func(t *testing.T) {

		auth.SetKey("testkey")
		if auth.GetKey() != "testkey" {
			t.Error("failed to set key")
		}

		auth.SetSecret("testsecret")
		if auth.GetSecret() != "testsecret" {
			t.Error("failed to set secret")
		}
	})
}
