package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/models"
)

func registerForEvent(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
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

	users, err := event.GetRegistrations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not get current registrations: %v", err),
		})
		return
	}

	for _, user := range users {
		if user == userId {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": "User already registered for this event",
			})
			return
		}
	}

	if err := event.Register(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not register for event: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registered",
	})
}
