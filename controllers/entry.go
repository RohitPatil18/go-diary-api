package controllers

import (
	"diary_api/helpers"
	"diary_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddEntry(context *gin.Context) {
	var entryIn models.Entry

	if err := context.ShouldBindJSON(&entryIn); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entryIn.UserID = user.ID

	newEntry, err := entryIn.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newEntry})
}

func GetAllEntries(context *gin.Context) {
	user, err := helpers.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Entries})
}
