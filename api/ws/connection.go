package ws

import (
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	maxMessageSize = 256 * 256
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  maxMessageSize,
	WriteBufferSize: maxMessageSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connection represents a single clients tile dispatcher.
type Connection struct {
	conn    *websocket.Conn
	handler RequestHandler
	mutex   sync.Mutex
}

// RequestHandler represents a handler for the ws request.
type RequestHandler func(*Connection, []byte)

// NewConnection returns a pointer to a new tile dispatcher object.
func NewConnection(w http.ResponseWriter, r *http.Request, handler RequestHandler) (*Connection, error) {
	// open a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	// set the message read limit
	conn.SetReadLimit(maxMessageSize)
	return &Connection{
		conn:    conn,
		handler: handler,
		mutex:   sync.Mutex{},
	}, nil
}

// ListenAndRespond waits on both tile request and responses and handles each
// until the websocket connection dies.
func (c *Connection) ListenAndRespond() error {
	for {
		// wait on read
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return err
		}
		// handle the message
		go c.handler(c, msg)
	}
}

// Send will send a json response in a thread safe manner.
func (c *Connection) Send(res interface{}) error {
	// writes are not thread safe
	c.mutex.Lock()
	defer runtime.Gosched()
	defer c.mutex.Unlock()
	// write response to websocket
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteJSON(res)
}

// Close closes the dispatchers websocket connection.
func (c *Connection) Close() {
	// close websocket connection
	c.conn.Close()
}
