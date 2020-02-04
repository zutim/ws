# ws
use websocket with gin

## Install
```
go get -u github.com/ebar-go/ws
```

## Usage
```go
package main
import (
    "github.com/ebar-go/ws"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    m := ws.New()

	// use gin router
	router := gin.Default()
	router.GET("/ws", func(ctx *gin.Context) {
		// get websocket conn
		conn, err := ws.GetUpgradeConnection(ctx.Writer, ctx.Request)
		if err != nil {
			http.NotFound(ctx.Writer, ctx.Request)
			return
		}
        
		connection := ws.NewConnection(conn, func(ctx *ws.Context) string {
			return ctx.GetMessage()
		})
        // register connection
		m.Register(connection)
        // listen
		go connection.Listen()
	})
    // start service
	go m.Start()

	_ = router.Run()    
}
```