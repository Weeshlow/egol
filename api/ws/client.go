package ws

import (
	"github.com/unchartedsoftware/egol/api/util"
)

// Client represents a single client connected via websockets
type Client struct {
	ID   string
	Conn *Connection
	New  bool
}

// NewClient returns a new client instance.
func NewClient() *Client {
	return &Client{
		ID:  util.RandID(),
		New: true,
	}
}
