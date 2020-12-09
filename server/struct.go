package server

import (
	"github.com/gin-contrib/cors"
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

// Here we create our new server
func NewServer() *Server {
	router := gin.New()
	// here we opened cors for all
	router.Use(cors.Default())
	return &Server{
		DB:     nil,
		Router: router,
	}
}
