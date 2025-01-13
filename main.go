package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zjunaidz/gin-rest-api/db"
	"github.com/zjunaidz/gin-rest-api/models"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()

	server.GET("/get", get)
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/create/event", createNewEvent)

	server.Run(":8080")
}

func get(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get events",
		})
		return
	}
	context.JSON(http.StatusOK, events)

}
func createNewEvent(context *gin.Context) {
	//* expect some data from client

	var event models.Event

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"message": "Could not parse data",
		})
		return
	}
	// * save event
	event.UserId = 1
	event.DateTime = time.Now()
	// * return event
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save Event",
			"error":err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse event id",
		})
		return
	}
	// * get event by id
	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get event",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Event found",
		"event":   event,
	})
}
