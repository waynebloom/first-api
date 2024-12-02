package routes

import (
	"golearn/first-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.GET("/users", getUsers)
	server.POST("/signup", postUser)
	server.POST("/login", login)

	auth := server.Group("/")
	auth.Use(middleware.Authenticate)
	auth.POST("/events", postEvent)
	auth.PUT("/events/:id", putEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.POST("/events/:id/register", registerForEvent)
	auth.GET("/events/:id/roster", getEventRoster)
	auth.DELETE("/events/:id/register", deregisterForEvent)
}
