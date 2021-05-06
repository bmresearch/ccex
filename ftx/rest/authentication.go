package rest

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"
)

// SignRequest is used to sign the request with the necessary information
func (c *Client) SignRequest(method, endpoint, url string, body []byte) *http.Request {
	ts := strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)

	signaturePayload := ts + method + endpoint + string(body)
	signature := c.Sign(signaturePayload)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("FTX-KEY", c.config.Auth.Key)
	req.Header.Set("FTX-SIGN", signature)
	req.Header.Set("FTX-TS", ts)

	if c.config.SubAccount != "" {
		req.Header.Set("FTX-SUBACCOUNT", c.config.SubAccount)
	}

	return req
}

// Sign is used to sign the payload for authentication.
func (c *Client) Sign(payload string) string {
	mac := hmac.New(sha256.New, []byte(c.config.Auth.Secret))
	mac.Write([]byte(payload))
	return hex.EncodeToString(mac.Sum(nil))
}
