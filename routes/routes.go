package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/middlewares"
)

func Register(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	authenticated.POST("/events/:id/register", createRegistration)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
