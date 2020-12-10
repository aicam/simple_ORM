package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

// Server parts : database(mysql) - router
type Server struct {
	DB               *gorm.DB
	Router           *gin.Engine
	SocketConnection websocket.Upgrader
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		c.Next()
	}
}

// Here we create our new server
func NewServer() *Server {
	router := gin.Default()
	// here we opened cors for all
	router.Use(CORS())
	return &Server{
		DB:     nil,
		Router: router,
	}
}
