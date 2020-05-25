package ws

// Handler define message processor
type Handler func(ctx *Context,c *Connection) string

// DefaultHandler
func DefaultHandler(ctx *Context,c *Connection) string {
	// do nothing
	return ctx.message
}
