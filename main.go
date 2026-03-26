package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/db"
	"github.com/vinicius-benevides/go-rest-api/models"
)

func main() {
	server := gin.Default()
	if err := db.Init(); err != nil {
		panic(err)
	}

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events := models.GetAllEvents()

	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Could not parse event data: %v", err),
		})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
