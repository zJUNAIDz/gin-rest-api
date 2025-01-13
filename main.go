package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zjunaidz/gin-rest-api/db"
	"github.com/zjunaidz/gin-rest-api/routes"
)

func main() {
	db.InitDB()
	defer db.DB.Close()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
