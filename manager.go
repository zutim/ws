package ws


// manager implement ws manager interface
type manager struct {
	connections    map[string]*Connection
	register   chan *Connection
	unregister chan *Connection
}

// Register register conn
func (manager *manager) Register(conn *Connection) {
	manager.register <- conn
}

// UnRegister delete websocket connection
func (manager *manager) Unregister(conn *Connection) {
	manager.unregister <- conn
}

// Start
func (manager *manager) Start() {
	for {
		select {
		case conn := <-manager.register:
			manager.connections[conn.ID] = conn

			conn.manager = manager
		case conn := <-manager.unregister:
			if _, ok := manager.connections[conn.ID]; ok {
				delete(manager.connections, conn.ID)
			}
		}
	}
}

// Broadcast push message to all connection, except ignore connection
func (manager *manager) Broadcast(message []byte, ignore *Connection) {
	for id, conn := range manager.connections {
		if ignore == nil || ignore.ID != id {
			conn.Send(message)
		}
	}
}

func (manager *manager)GetOnLine() int{
	return len(manager.connections)
}