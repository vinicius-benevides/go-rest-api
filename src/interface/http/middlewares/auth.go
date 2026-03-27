package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/src/services"
)

func Authenticate(tokenService services.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Could not authenticate user",
			})
			return
		}

		userId, err := tokenService.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Could not authenticate user",
			})
			return
		}

		ctx.Set("userId", userId)
		ctx.Next()
	}
}
