package ws

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	m := New()
	if m == nil {
		t.Error("manager is nil")
	}
}

func TestManager_Start(t *testing.T) {
	m := New()

	// use gin router
	router := gin.Default()
	router.GET("/ws", func(ctx *gin.Context) {
		// get websocket conn
		conn, err := GetUpgradeConnection(ctx.Writer, ctx.Request)
		if err != nil {
			http.NotFound(ctx.Writer, ctx.Request)
			return
		}

		connection := NewConnection(conn, func(ctx *Context) string {
			return ctx.GetMessage()
		})

		m.Register(connection)

		go connection.Listen()
	})
	go m.Start()

	_ = router.Run()
}
