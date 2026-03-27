package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
)

func RespondError(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	message := "Internal server error"

	var appErr *errs.AppError
	if errors.As(err, &appErr) {
		message = appErr.Error()
		switch appErr.Code() {
		case errs.NotFound:
			status = http.StatusNotFound
		case errs.Conflict:
			status = http.StatusConflict
		case errs.Unauthorized:
			status = http.StatusUnauthorized
		case errs.BadRequest:
			status = http.StatusBadRequest
		default:
			status = http.StatusInternalServerError
		}
	}

	ctx.JSON(status, gin.H{
		"message": message,
	})
}
