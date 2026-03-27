package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/pkg/errs"
	"github.com/vinicius-benevides/go-rest-api/src/interface/http/helpers"
	"github.com/vinicius-benevides/go-rest-api/src/models"
	"github.com/vinicius-benevides/go-rest-api/src/services"
)

var userService = services.NewUserService()

func signup(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.RespondError(ctx, errs.BadRequestError("Invalid user payload"))
		return
	}

	createdUser, err := userService.Signup(user)
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"user":    createdUser,
	})
}

func login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		helpers.RespondError(ctx, errs.BadRequestError("Invalid credentials payload"))
		return
	}

	token, err := userService.Login(user.Email, user.Password)
	if err != nil {
		helpers.RespondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
		"token":   token,
	})
}
