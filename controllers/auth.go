package controllers

import (
	"diary_api/helpers"
	"diary_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var authIn models.AuthenticationIn

	if err := context.ShouldBindJSON(&authIn); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: authIn.Username,
		Password: authIn.Password,
	}

	newUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func Login(context *gin.Context) {
	var authIn models.AuthenticationIn

	if err := context.ShouldBindJSON(&authIn); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(authIn.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(authIn.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := helpers.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessToken": accessToken})
}
