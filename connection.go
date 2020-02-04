package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var u = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} // use default options

// GetUpgradeConnection get web socket connection
func GetUpgradeConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	respHeader := http.Header{"Sec-WebSocket-Protocol": []string{r.Header.Get("Sec-WebSocket-Protocol")}}
	return u.Upgrade(w, r, respHeader)
}

// Connection include socket conn
type Connection struct {
	// unique id
	ID string

	// socket connection
	conn *websocket.Conn

	// process handler
	Handler Handler

	// connection manager
	manager Manager
}


// Send push message to client
func (c *Connection) Send(message []byte) {
	_ = c.conn.WriteMessage(websocket.TextMessage, message)
}

// close
func (c *Connection) close() {
	_ = c.conn.Close()
	c.manager.Unregister(c)
}

// Listen listen connection
func (c *Connection) Listen() {
	defer func() {
		c.close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		ctx := &Context{message: string(message)}

		result := c.Handler(ctx)
		c.Send([]byte(result))
	}
}
