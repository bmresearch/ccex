package rest

import (
	"github.com/valyala/fasthttp"
)

// Client is the base http client
type Client struct {
	BaseUrl  string
	Request  *fasthttp.Request
	Response *fasthttp.Response
	client   fasthttp.Client
}

// Submit submits an HTTP request.
func (c *Client) Submit() error {
	err := c.client.Do(c.Request, c.Response)
	return err
}

// New returns a new client for HTTP requests.
func New(url string) (*Client, error) {
	client := &Client{
		client:   fasthttp.Client{},
		BaseUrl:  url,
		Request:  fasthttp.AcquireRequest(),
		Response: fasthttp.AcquireResponse(),
	}

	return client, nil
}
