package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/murlokito/ccex/auth"
	"golang.org/x/time/rate"

	"github.com/murlokito/ccex/config"
	"github.com/murlokito/ccex/internal/rest"
)

// SignatureFunc is the
type SignatureFunc func(method string, path string, body []byte) *http.Request

// Client represents the REST API Client for FTX.
type Client struct {
	config  *config.Configuration
	client  *rest.Client
	limiter *rate.Limiter

	auth *auth.Authentication
}

/*
	Get fetches the information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Get(endpoint string) ([]byte, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ErrRateLimited, duration.Milliseconds())
	}

	reqUrl := c.client.BaseUrl + endpoint

	preparedRequest := c.SignRequest("GET", endpoint, reqUrl, []byte(""))

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()

	var buffer []byte

	_, err = resp.Body.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

/*
	Post submits information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Post(endpoint string, data map[string]interface{}) ([]byte, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ErrRateLimited, duration.Milliseconds())
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqUrl := c.client.BaseUrl + endpoint

	preparedRequest := c.SignRequest("POST", endpoint, reqUrl, payload)

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()

	var buffer []byte

	_, err = resp.Body.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

/*
	Put submits information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Put(endpoint string, data map[string]interface{}) ([]byte, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ErrRateLimited, duration.Milliseconds())
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqUrl := c.client.BaseUrl + endpoint

	preparedRequest := c.SignRequest("PUT", endpoint, reqUrl, payload)

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()

	var buffer []byte

	_, err = resp.Body.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

/*
	Put submits information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Delete(endpoint string, data map[string]interface{}) ([]byte, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ErrRateLimited, duration.Milliseconds())
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	reqUrl := c.client.BaseUrl + endpoint

	preparedRequest := c.SignRequest("DELETE", endpoint, reqUrl, payload)

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()

	var buffer []byte

	_, err = resp.Body.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// NewClient returns a new rest client for ftx
func NewClient(cfg *config.Configuration) (*Client, error) {
	rc, err := rest.New(cfg, ApiUrl)
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		client:  rc,
		limiter: rate.NewLimiter(30, 5),
	}

	return client, nil
}
