package ws

// Handler define message processor
type Handler func(ctx *Context) string

// DefaultHandler
func DefaultHandler(ctx *Context) string {
	// do nothing
	return ctx.message
}
