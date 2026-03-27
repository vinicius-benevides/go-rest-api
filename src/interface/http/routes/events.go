package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/interface/http/helpers"
	"github.com/vinicius-benevides/go-rest-api/src/models"
	"github.com/vinicius-benevides/go-rest-api/src/services"
)

var eventService = services.NewEventService()

func getEvents(ctx *gin.Context) {
	events, err := eventService.ListEvents()
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func getEventByID(ctx *gin.Context) {
	eventId, err := helpers.GetIDParam(ctx, "id", "event id")
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	event, err := eventService.GetEvent(eventId)
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func createEvent(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		helpers.RespondError(ctx, errs.BadRequestError("Invalid event payload"))
		return
	}

	userId := ctx.GetInt64("userId")
	created, err := eventService.CreateEvent(userId, event)
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   created,
	})
}

func updateEvent(ctx *gin.Context) {
	eventId, err := helpers.GetIDParam(ctx, "id", "event id")
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	var updatedEvent models.Event
	if err := ctx.ShouldBindJSON(&updatedEvent); err != nil {
		helpers.RespondError(ctx, errs.BadRequestError("Invalid event payload"))
		return
	}

	userId := ctx.GetInt64("userId")
	event, err := eventService.UpdateEvent(userId, eventId, updatedEvent)
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   event,
	})
}

func deleteEvent(ctx *gin.Context) {
	eventId, err := helpers.GetIDParam(ctx, "id", "event id")
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	userId := ctx.GetInt64("userId")
	if err := eventService.DeleteEvent(userId, eventId); err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event deleted",
	})
}
