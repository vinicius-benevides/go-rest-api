package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/src/interface/http/helpers"
	"github.com/vinicius-benevides/go-rest-api/src/services"
)

var registrationService = services.NewRegistrationService()

func createRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := helpers.GetIDParam(ctx, "id", "event id")
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	if err := registrationService.RegisterUser(eventId, userId); err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registered",
	})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := helpers.GetIDParam(ctx, "id", "event id")
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	if err := registrationService.CancelRegistration(eventId, userId); err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Canceled",
	})
}
