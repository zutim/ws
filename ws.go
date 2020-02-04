package ws

import (
	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid")

// Manager websocket manager
type Manager interface {
	// register connection
	Register(conn *Connection)

	// unregister connection
	Unregister(conn *Connection)

	// broadcast message
	Broadcast(message []byte, ignore *Connection)

	// start service
	Start()
}

// New return ws manager instance
func New() Manager {
	return &manager{
		connections:    make(map[string]*Connection),
		register:   make(chan *Connection),
		unregister: make(chan *Connection),
	}
}


// NewConnection return Connection
func NewConnection(conn *websocket.Conn, handler Handler) *Connection {
	if handler == nil {
		handler = DefaultHandler
	}
	return &Connection{
		ID:      uuid.NewV4().String(),
		conn:    conn,
		Handler: handler,
	}
}