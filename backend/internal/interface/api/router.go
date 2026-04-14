package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"

	"github.com/the-sandwich/backend/internal/interface/ws"
)

func SetupRouter(hub *ws.Hub, handlers *Handlers) (*gin.Engine, *http.Server) {
	r := gin.Default()

	// CORS config
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/register", handlers.RegisterHandler)
		api.POST("/login", handlers.LoginHandler)
	}

	// Socket.IO
	r.GET("/socket.io/:any", gin.WrapH(hub.Server))
	r.POST("/socket.io/:any", gin.WrapH(hub.Server))

	addr := ":8080"
	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return r, server
}

func NewSocketIOHandler(server *socketio.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		server.ServeHTTP(c.Writer, c.Request)
		if c.Request.Method == "GET" {
			c.AbortWithStatus(http.StatusUpgradeRequired)
		}
	}
}
