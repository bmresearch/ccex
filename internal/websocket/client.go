package websocket

import (
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/murlokito/ccex/log"

	"github.com/gorilla/websocket"
	"github.com/jpillora/backoff"
	"github.com/pkg/errors"
)

// OnMessageHandler is a handler that is dispatched when the client reads a message from the websocket.
type OnMessageHandler func(message []byte) error

/*
OnConnectedHandler is a handler that is dispatched when the client reads a message from the websocket.
This can be the authentication message and/or something else.
*/
type OnConnectedHandler func() ([]byte, error)

// ErrNotConnected is returned when the application read/writes
// a message and the connection is closed
var ErrNotConnected = errors.New("websocket: not connected")

// Client type represents a Reconnecting WebSocket connection.
type Client struct {
	// RecIntvlMin specifies the initial reconnecting interval,
	// default to 2 seconds.
	RecIntvlMin time.Duration

	// RecIntvlMax specifies the maximum reconnecting interval,
	// default to 30 seconds.
	RecIntvlMax time.Duration

	// RecIntvlFactor specifies the rate of increase of the reconnection
	// interval, default to 1.5.
	RecIntvlFactor float64

	// Backoff holds the previous attributes in order to allow the reconnection
	// mechanism to work.
	Backoff *backoff.Backoff

	// HandshakeTimeout specifies the duration for the handshake to complete,
	// default to 2 seconds.
	HandshakeTimeout time.Duration

	// Verbose suppress connecting/reconnecting messages.
	Verbose bool

	// Conn represents the underlying websocket connection.
	Conn *websocket.Conn

	// Dialer represents the mechanism used to perform the connection.
	Dialer *websocket.Dialer

	// OnConnected represents the handler dispatched when the connection is opened.
	OnConnected OnConnectedHandler

	// OnMessage represents the handler dispatched when a message is read.
	OnMessage OnMessageHandler

	// logger holds the logger to be used.
	logger log.Logger

	// mu is the mutex to prevent issues associated with goroutine concurrency.
	mu sync.RWMutex

	// reqHeader holds the http request header to be used during connection.
	reqHeader http.Header

	// httpResp holds the http response that is received from the connection
	httpResp *http.Response

	// keepAliveTimeout is an interval for sending ping/pong messages disabled if 0.
	keepAliveTimeout time.Duration

	// url holds the url for the connection.
	url string

	// dialErr is used to hold a possible error caught when dialing the server.
	dialErr error

	// verbose is used to define if the client should log actions verbosely and/or log more actions.
	verbose bool

	// connected is used to define if the client has an established connection.
	connected bool

	// connected is used to define if the client's connection is closed.
	closed bool
}

// CloseAndReconnect will try to reconnect.
func (c *Client) CloseAndReconnect() (err error) {
	err = c.Close()
	if err != nil {
		return err
	}

	go c.Connect()
	return nil
}

// GetBackoff retrieves the backoff associated with the connection.
func (c *Client) GetBackoff() *backoff.Backoff {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return &backoff.Backoff{
		Min:    c.RecIntvlMin,
		Max:    c.RecIntvlMax,
		Factor: c.RecIntvlFactor,
		Jitter: true,
	}
}

// SetKeepAliveTimeout sets the interval for the keep alive message
func (c *Client) SetKeepAliveTimeout(interval time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.keepAliveTimeout = interval
}

// Connect performs the connection. This function should be executed by a goroutine.
func (c *Client) Connect() {
	b := c.GetBackoff()
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		nextReconnect := b.Duration()
		wsConn, httpResp, err := c.Dialer.Dial(c.url, c.reqHeader)

		c.mu.Lock()
		c.Conn = wsConn
		c.dialErr = err
		c.connected = err == nil
		c.httpResp = httpResp
		c.mu.Unlock()

		if c.Closed() {
			c.logger.Info("Something went wrong, retrying.")
			if err := c.CloseAndReconnect(); err != nil {
				c.logger.Error(err.Error())
			}
			return
		}

		if !c.GetVerbose() && err != nil {
			c.logger.Error(err.Error())
			c.logger.Info("Retrying connection.")
			if err := c.CloseAndReconnect(); err != nil {
				c.logger.Error(err.Error())
			}
		}

		if err == nil {
			if !c.GetVerbose() {
				c.logger.Info("Connection was successfully established with ", c.url)

				if c.OnConnected != nil {
					message, err := c.OnConnected()
					if err == nil {
						err = c.WriteMessage(1, message)
						if c.getVerbose() && err != nil {
							c.logger.Error(err.Error())

							// In case connection ended for some reason.
							if c.Closed() {
								c.logger.Info("Something went wrong, retrying.")
								if err := c.CloseAndReconnect(); err != nil {
									c.logger.Error(err.Error())
								}
								return
							}
						}
					}
				}

				if c.getKeepAliveTimeout() != 0 {
					c.logger.Info("Keeping connection alive with timeout ", c.getKeepAliveTimeout())
					c.keepAlive()
				}

				err = c.ReadMessages()
				if err != nil {
					c.logger.Error(err.Error())
				}

				if err == ErrNotConnected {
					if err := c.CloseAndReconnect(); err != nil {
						c.logger.Error(err.Error())
					}
					return
				}
			}

			return
		}

		if c.getVerbose() {
			c.logger.Error(err.Error())
			c.logger.Info("Will try again in ", nextReconnect, " seconds.")
		}

		time.Sleep(nextReconnect)
	}
}

// Connected returns the WebSocket connection state
func (c *Client) Connected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.connected
}

// Closed returns the WebSocket connection state
func (c *Client) Closed() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.closed
}

// setConnected sets state for Connected
func (c *Client) setConnected(state bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.connected = state
}

// setClosed sets state for Closed
func (c *Client) setClosed(state bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.closed = state
}

// getConn gets the underlying connection.
func (c *Client) getConn() *websocket.Conn {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Conn
}

// Close closes the underlying network connection without
// sending or waiting for a close frame.
func (c *Client) Close() error {
	c.setClosed(true)

	if c.getConn() != nil {
		c.mu.Lock()
		err := c.Conn.Close()
		if err != nil {
			return err
		}
		c.mu.Unlock()
	}

	c.setConnected(false)
	return nil
}

// ReadMessages reads messages while the connection is active
func (c *Client) ReadMessages() error {
	for {
		if c.Closed() {
			return ErrNotConnected
		}

		_, msg, err := c.ReadMessage()
		if err != nil {
			return err
		}

		err = c.OnMessage(msg)
		if err != nil {
			c.logger.Error(err.Error())
		}
	}
}

// ReadMessage is a helper method for reading a message from the underlying connection.
// If the connection is closed ErrNotConnected is returned
func (c *Client) ReadMessage() (messageType int, message []byte, err error) {
	if !c.Connected() {
		return 0, nil, ErrNotConnected
	}

	messageType, message, err = c.Conn.ReadMessage()
	if err != nil {
		return 0, nil, err
	}

	return 0, message, nil
}

// WriteMessage is a helper method for writing a message to the underlying connection.
// If the connection is closed ErrNotConnected is returned
func (c *Client) WriteMessage(messageType int, data []byte) (err error) {
	if !c.Connected() {
		return ErrNotConnected
	}

	c.mu.Lock()
	err = c.Conn.WriteMessage(messageType, data)
	c.mu.Unlock()
	if err != nil {
		return err
	}

	return err
}

// WriteJSON writes the JSON encoding of v to the connection.
//
// See the documentation for encoding/json Marshal for details about the
// conversion of Go values to JSON.
//
// If the connection is closed ErrNotConnected is returned
func (c *Client) WriteJSON(v interface{}) (err error) {
	if !c.Connected() {
		return ErrNotConnected
	}

	c.mu.Lock()
	err = c.Conn.WriteJSON(v)
	c.mu.Unlock()
	if err != nil {
		return err
	}

	return err
}

// ReadJSON reads the next JSON-encoded message from the connection and stores
// it in the value pointed to by v.
//
// See the documentation for the encoding/json Unmarshal function for details
// about the conversion of JSON to a Go value.
//
// If the connection is closed ErrNotConnected is returned
func (c *Client) ReadJSON(v interface{}) (err error) {
	if !c.Connected() {
		return ErrNotConnected
	}

	if c.Connected() {
		err = c.Conn.ReadJSON(v)
		if err != nil {
			return err
		}
	}

	return err
}

// validateUrl validates passed rawUrl.
func validateUrl(rawUrl string) (string, error) {
	if rawUrl == "" {
		return "", errors.New("dial: url cannot be empty")
	}

	u, err := url.Parse(rawUrl)

	if err != nil {
		return "", errors.New("url: " + err.Error())
	}

	if u.Scheme != "ws" && u.Scheme != "wss" {
		return "", errors.New("url: websocket uris must start with ws or wss scheme")
	}

	if u.User != nil {
		return "", errors.New("url: user name and password are not allowed in websocket URIs")
	}

	return rawUrl, nil
}

// setDefaultRecIntvlMin sets the default reconnection interval minimum time.
func (c *Client) setDefaultRecIntvlMin() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.RecIntvlMin == 0 {
		c.RecIntvlMin = 1 * time.Second
	}
}

// setDefaultRecIntvlMax sets the default reconnection interval maximum time.
func (c *Client) setDefaultRecIntvlMax() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.RecIntvlMax == 0 {
		c.RecIntvlMax = 5 * time.Second
	}
}

// setDefaultRecIntvlFactor sets the default reconnection interval factor.
func (c *Client) setDefaultRecIntvlFactor() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.RecIntvlFactor == 0 {
		c.RecIntvlFactor = 1.5
	}
}

// setDefaultHandshakeTimeout sets the default handshake timeout.
func (c *Client) setDefaultHandshakeTimeout() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.HandshakeTimeout == 0 {
		c.HandshakeTimeout = 2 * time.Second
	}
}

// setDefaultDialer sets the default dialer
func (c *Client) setDefaultDialer(handshakeTimeout time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Dialer = &websocket.Dialer{
		HandshakeTimeout: handshakeTimeout,
	}
}

// getHandshakeTimeout returns the duration to wait for handshake to complete.
func (c *Client) getHandshakeTimeout() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.HandshakeTimeout
}

// Dial establishes a new client connection.
// The URL url specifies the host and request URI. Use requestHeader to specify
// the origin (Origin), sub-protocols (Sec-WebSocket-Protocol) and cookies
// (Cookie).
func (c *Client) Dial() {

	// Config
	c.setDefaultRecIntvlMin()
	c.setDefaultRecIntvlMax()
	c.setDefaultRecIntvlFactor()
	c.setDefaultHandshakeTimeout()
	c.setDefaultDialer(c.getHandshakeTimeout())

	// Connect
	go c.Connect()

	// wait on first attempt
	time.Sleep(c.getHandshakeTimeout())
}

// GetVerbose returns `Verbose`, a boolean used to define if the client should suppress connection/reconnection messages or not. False by default.
func (c *Client) GetVerbose() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Verbose
}

// getVerbose returns `verbose`, a boolean used to define if the client should suppress logging messages or not. True by default.
func (c *Client) getVerbose() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.verbose
}

// getKeepAliveTimeout returns `keepAliveTimeout` which defines the time between Ping messages.
func (c *Client) getKeepAliveTimeout() time.Duration {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.keepAliveTimeout
}

// writeControlPingMessage writes a Ping Control message.
func (c *Client) writeControlPingMessage() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.Conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second))
}

// keepAlive is used to keep the connection alive, it should be executed in a goroutine.
func (c *Client) keepAlive() {
	var (
		keepAliveResponse = new(keepAliveResponse)
		ticker            = time.NewTicker(c.getKeepAliveTimeout())
	)

	c.mu.Lock()
	c.Conn.SetPongHandler(func(msg string) error {
		keepAliveResponse.setLastResponse()
		return nil
	})
	c.mu.Unlock()

	go func() {
		defer ticker.Stop()

		for {
			if c.getVerbose() {
				c.logger.Info("Writing ping message")
			}

			if err := c.writeControlPingMessage(); err != nil {
				c.logger.Error(err.Error())
			}
			<-ticker.C
			if c.Closed() {
				return
			}
			if time.Now().Sub(keepAliveResponse.getLastResponse()) > c.getKeepAliveTimeout() {
				if err := c.CloseAndReconnect(); err != nil {
					c.logger.Error(err.Error())
				}
				return
			}
		}
	}()
}

// New returns a new configured websocket client for the passed url.
func New(url string, logger log.Logger) (*Client, error) {

	validatedUrl, err := validateUrl(url)
	if err != nil {
		return nil, err
	}

	client := &Client{
		url:    validatedUrl,
		logger: logger,
	}

	return client, nil
}
