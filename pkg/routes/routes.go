package routes

import (
	"apiv2/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

// pointer to our Gin app
func RegisterRoutes(server *gin.Engine) {
	// create a router group
	authenticated := server.Group("/")
	// use auth middleware, ensure it's before
	authenticated.Use(middlewares.Authenticated)
	// url and handler
	server.GET("/events", GetEvents)
	// get event by ID
	server.GET("/events/:id", GetEvent)
	authenticated.POST("/events", CreateEvents)
	//put to update resource
	authenticated.PUT("events/:id", UpdateEvent)
	authenticated.DELETE("events/:id", DeleteEvent)
	server.POST("/signup", CreateUsers)
	server.POST("/login", Login)
}
