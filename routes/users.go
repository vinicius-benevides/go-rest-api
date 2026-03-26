package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/models"
)

func signup(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Could not parse user data: %v", err),
		})
		return
	}

	if existingUser, err := models.GetUserByEmail(user.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not verify user: %v", err),
		})
		return
	} else if existingUser != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "Email already registered",
		})
		return
	}

	if err := user.Save(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Could not create user: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"user":    user,
	})
}

func login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Could not parse user data: %v", err),
		})
		return
	}

	if err := user.Login(); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": fmt.Sprintf("Could not authenticate user: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login Successful",
	})
}
