package rest

import (
	"encoding/json"
	"fmt"
	"github.com/murlokito/ccex/auth"
	"github.com/murlokito/ccex/ftx"
	"golang.org/x/time/rate"
	"net/http"
	"time"

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
func (c *Client) Get() (*http.Response, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ftx.ErrRateLimited, duration.Milliseconds())
	}

	preparedRequest := c.SignRequest("GET", c.client.BaseUrl, []byte(""))

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()
	return resp, nil
}

/*
	Post submits information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Post(data map[string]interface{}) (*http.Response, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ftx.ErrRateLimited, duration.Milliseconds())
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	preparedRequest := c.SignRequest("POST", c.client.BaseUrl, payload)

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()
	return resp, nil
}

/*
	Put submits information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Put(data map[string]interface{}) (*http.Response, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ftx.ErrRateLimited, duration.Milliseconds())
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	preparedRequest := c.SignRequest("PUT", c.client.BaseUrl, payload)

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()
	return resp, nil
}

/*
	Put submits information.

	This method is only implemented in every exchange in case some specific
	action must be taken before processing the request, which generally it does,
	due to authentication, etc.
*/
func (c *Client) Delete(data map[string]interface{}) (*http.Response, error) {
	reservation := c.limiter.Reserve()

	if !reservation.OK() {
		duration := reservation.DelayFrom(time.Now())
		reservation.Cancel()

		return nil, fmt.Errorf(ftx.ErrRateLimited, duration.Milliseconds())
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	preparedRequest := c.SignRequest("DELETE", c.client.BaseUrl, payload)

	resp, err := c.client.Submit(preparedRequest)
	if err != nil {
		return nil, err
	}

	reservation.Cancel()
	return resp, nil
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
