package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	server.GET("/events/:id", getEventByID)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not get events: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventByID(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Could not parse event id: %v", err),
		})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not get event: %v", err),
		})
		return
	}

	if event == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, event)
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

	if err := event.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not create event: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
