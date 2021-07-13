package rest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// PrepareRequest is used to prepare the request with the necessary information, including authentication info
func (c *Client) PrepareRequest(method, endpoint, endpointUrl string, extraParams map[string]interface{}) error {
	c.client.Request.Header.SetMethod(method)

	if c.config != nil {
		err := c.PrepareSignedRequest(method, endpoint, endpointUrl, extraParams)
		if err != nil {
			return err
		}
	} else {
		orderedParams := GetOrderedParameters(extraParams)
		url := fmt.Sprintf("%s%s", endpointUrl, orderedParams)
		fmt.Println(url)
		c.client.Request.SetRequestURI(url)
	}

	return nil
}

// PrepareSignedRequest is used to fill the authentication data needed for authentication.
func (c *Client) PrepareSignedRequest(method, endpoint, endpointUrl string, extraParams map[string]interface{}) error {
	authParams := map[string]interface{}{
		"api_key":   c.config.Auth.Key,
		"timestamp": strconv.FormatInt(time.Now().UTC().Unix()*1000, 10),
	}

	// Merge the authentication parameters with the query parameters
	for k, v := range extraParams {
		authParams[k] = v
	}

	orderedParams := GetOrderedParameters(authParams)

	// Get the signature of the data which includes the API key, timestamp and extra params
	signature, err := c.Sign(orderedParams)
	if err != nil {
		return err
	}

	switch method {
	case http.MethodGet:
		// Append the signature to the parameters string and set the reque
		orderedParams += "&sign=" + signature
		url := fmt.Sprintf("%s%s", endpointUrl, orderedParams)
		fmt.Println(url)
		c.client.Request.SetRequestURI(url)
		break
	default:
		// Add the signature to the payload and marshal for request body
		authParams["sign"] = signature
		c.client.Request.SetRequestURI(endpointUrl)
		fmt.Println("url: " + endpointUrl + "payload: " + fmt.Sprintf("%v", authParams))
		payload, err := json.Marshal(authParams)
		if err != nil {
			return err
		}
		c.client.Request.SetBody(payload)
		break
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

// GetOrderedParameters is used to build a string with the passed parameters ordered alphabetically.
func GetOrderedParameters(params map[string]interface{}) string {
	var (
		keys          []string
		i             int
		orderedParams string
	)

	keys = make([]string, len(params))

	for k, _ := range params {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for j, k := range keys {
		if j+1 != len(keys) {
			orderedParams += k + "=" + fmt.Sprintf("%v", params[k]) + "&"
		}
		orderedParams += k + "=" + fmt.Sprintf("%v", params[k])
	}

	return orderedParams
}
