package rest

import (
	"net/http"

	"github.com/murlokito/ccex/config"
)

type Client struct {
	http.Client
	BaseUrl string
}

// Submit submits an HTTP request.
func (c *Client) Submit(request *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(request)
	return resp, err
}

// New returns a new client for HTTP requests.
func New(cfg *config.Configuration, url string) (*Client, error) {
	client := &Client{
		BaseUrl: url,
	}

	return client, nil
}
