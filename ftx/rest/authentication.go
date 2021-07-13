package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

// PrepareRequest is used to prepare the request with the necessary information, including authentication info
func (c *Client) PrepareRequest(method, endpoint, url string, body []byte) error {

	c.client.Request.Header.SetMethod(method)

	c.client.Request.SetRequestURI(url)
	c.client.Request.SetBody(body)

	if c.config != nil {
		if err := c.SignRequest(method, endpoint, url, body); err != nil {
			return err
		}
	}
	return nil
}

// SignRequest is used to fill the authentication data
func (c *Client) SignRequest(method, endpoint, url string, body []byte) error {
	var (
		signature string
		ts        string
	)

	ts = strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)

	signaturePayload := ts + method + endpoint + string(body)
	signature, err := c.Sign(signaturePayload)
	if err != nil {
		return err
	}

	c.client.Request.Header.Set("FTX-KEY", c.config.Auth.Key)
	c.client.Request.Header.Set("FTX-SIGN", signature)
	c.client.Request.Header.Set("FTX-TS", ts)

	if c.config.SubAccount != "" {
		c.client.Request.Header.Set("FTX-SUBACCOUNT", c.config.SubAccount)
	}

	return nil
}

// Sign is used to sign the payload for authentication.
func (c *Client) Sign(payload string) (string, error) {
	mac := hmac.New(sha256.New, []byte(c.config.Auth.Secret))
	_, err := mac.Write([]byte(payload))
	if err != nil {
		return "", errors.Wrap(err, "error calculating hmac for signature")
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}
