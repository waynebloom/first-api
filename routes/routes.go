package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", postEvent)
  server.PUT("/events/:id", putEvent)
  server.DELETE("events/:id", deleteEvent)
  server.GET("/users", getUsers)
  server.POST("/signup", postUser)
}
