package api

import (
	"github.com/gin-gonic/gin"
	"github.com/the-sandwich/backend/internal/ws"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS config (simplistic)
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
		api.POST("/register", RegisterHandler)
		api.POST("/login", LoginHandler)
	}

	r.GET("/ws", ws.ServeWS)

	return r
}
