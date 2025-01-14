package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zjunaidz/gin-rest-api/utils"
)

func Authentication(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized",
		})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized. Invalid token",
			"error":   err.Error(),
		})
		return
	}
	// Set the userId in the context
	context.Set("userId", userId)
	context.Next()
}
