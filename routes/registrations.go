package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zjunaidz/gin-rest-api/models"
)

func registerForEvent(context *gin.Context) {
	//* get event_id and user_id
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse event id to int64",
			"error":   err.Error(),
		})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get event",
			"error":   err.Error(),
		})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register for event",
			"error":   err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Registration Successful",
	})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse event id to int64",
			"error":   err.Error(),
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get event",
			"error":   err.Error(),
		})
		return
	}
	err = event.Unregister(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to unregister for event",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event Registration Cancelled",
	})
}
