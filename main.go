package main

import (
	"net/http"
	"sum/gin-api/db"
	"sum/gin-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/get", handelGet)
	server.GET("/get/events", handleGetEvents)
	server.POST("/create/event", handleCreateEvent)

	server.Run(":8080")
}

func handelGet(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}
func handleGetEvents(context *gin.Context) {
	events := models.GetAllEvent()
	context.JSON(http.StatusOK, events)

}
func handleCreateEvent(context *gin.Context) {
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
	event.Id = 1
	event.UserId = 1
	event.DateTime = time.Now()
	// * return event
	event.Save()

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}
