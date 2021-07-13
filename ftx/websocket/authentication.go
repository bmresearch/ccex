package websocket

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// AuthenticationMessage is used to build the websocket authentication message
func (c *Client) AuthenticationMessage() (string, string) {
	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)

	payload := ts + "websocket_login"

	mac := hmac.New(sha256.New, []byte(c.config.Auth.Secret))
	mac.Write([]byte(payload))

	return ts, hex.EncodeToString(mac.Sum(nil))
}
