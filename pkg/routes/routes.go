package routes

import (
	"apiv2/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// pointer to our Gin app
func RegisterRoutes(server *gin.Engine) {
	// url and handler
	server.GET("/events", GetEvents)
	// get event by ID
	server.GET("/events/:id", GetEvent)
	server.POST("/events", middlewares.Authenticated, CreateEvents)
	//put to update resource
	server.PUT("events/:id", UpdateEvent)
	server.DELETE("events/:id", DeleteEvent)
	server.POST("/signup", CreateUsers)
	server.POST("/login", Login)
}
