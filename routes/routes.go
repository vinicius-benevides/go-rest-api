package routes

import "github.com/gin-gonic/gin"

func Register(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvent)
}
