package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Register all routes here
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventById)
	server.PUT("/events/update/:id", updateEventById)
	server.POST("/events/create", createNewEvent)
	server.DELETE("/events/delete/:id", deleteEventById)
	server.GET("/users", getAllUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
