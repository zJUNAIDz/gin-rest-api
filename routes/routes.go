package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zjunaidz/gin-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// Register all routes here
	//*protected routes group
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authentication)
	authenticated.PUT("/events/update/:id", updateEventById)
	authenticated.POST("/events/create", createNewEvent)
	authenticated.DELETE("/events/delete/:id", deleteEventById)
	authenticated.POST("/events/register/:id", registerForEvent)
	authenticated.DELETE("/events/cancel/:id", cancelRegistration)
	server.GET("/events", getAllEvents)
	server.GET("/events/:id", getEventById)
	server.GET("/users", getAllUsers)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
