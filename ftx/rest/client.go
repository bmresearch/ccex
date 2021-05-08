package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/murlokito/ccex/auth"
	"github.com/pkg/errors"
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
	err := c.PrepareRequest(http.MethodGet, endpoint, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	err = c.client.Submit()
	if err != nil {
		return nil, errors.Wrap(err, "error submitting request")
	}

	reservation.Cancel()

	return c.client.Response.Body(), nil
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
	err = c.SignRequest(http.MethodPost, endpoint, reqUrl, payload)
	if err != nil {
		return nil, err
	}

	err = c.client.Submit()
	if err != nil {
		return nil, errors.Wrap(err, "error submitting request")
	}

	reservation.Cancel()

	return c.client.Response.Body(), nil
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
	err = c.SignRequest(http.MethodPut, endpoint, reqUrl, payload)
	if err != nil {
		return nil, err
	}

	err = c.client.Submit()
	if err != nil {
		return nil, errors.Wrap(err, "error submitting request")
	}

	reservation.Cancel()

	return c.client.Response.Body(), nil
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
	err = c.SignRequest(http.MethodDelete, endpoint, reqUrl, payload)
	if err != nil {
		return nil, err
	}

	err = c.client.Submit()
	if err != nil {
		return nil, errors.Wrap(err, "error submitting request")
	}

	reservation.Cancel()

	return c.client.Response.Body(), nil
}

// NewClient returns a new rest client for ftx
func NewClient(cfg *config.Configuration) (*Client, error) {
	rc, err := rest.New(ApiUrl)
	if err != nil {
		return &Client{}, err
	}

	client := &Client{
		config:  cfg,
		client:  rc,
		limiter: rate.NewLimiter(30, 5),
	}

	return client, nil
}
