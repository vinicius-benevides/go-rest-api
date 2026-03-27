package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/models"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/services"
)

var eventService = services.NewEventService()

func getEvents(ctx *gin.Context) {
	events, err := eventService.ListEvents()
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventByID(ctx *gin.Context) {
	eventId, err := getIDParam(ctx, "id", "event id")
	if err != nil {
		respondError(ctx, err)
		return
	}

	event, err := eventService.GetEvent(eventId)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		respondError(ctx, errs.BadRequestError("Invalid event payload"))
		return
	}

	userId := ctx.GetInt64("userId")
	created, err := eventService.CreateEvent(userId, event)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   created,
	})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := getIDParam(ctx, "id", "event id")
	if err != nil {
		respondError(ctx, err)
		return
	}

	var updatedEvent models.Event
	if err := ctx.ShouldBindJSON(&updatedEvent); err != nil {
		respondError(ctx, errs.BadRequestError("Invalid event payload"))
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := eventService.UpdateEvent(userId, eventId, updatedEvent)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   event,
	})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := getIDParam(ctx, "id", "event id")
	if err != nil {
		respondError(ctx, err)
		return
	}

	userId := ctx.GetInt64("userId")
	if err := eventService.DeleteEvent(userId, eventId); err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})
}
