package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
)

func getIDParam(ctx *gin.Context, paramName, label string) (int64, error) {
	value, err := strconv.ParseInt(ctx.Param(paramName), 10, 64)
	if err != nil {
		if label == "" {
			label = paramName
		}
		return 0, errs.BadRequestError("Invalid " + label)
	}

	return value, nil
}
