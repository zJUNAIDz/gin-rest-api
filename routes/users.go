package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zjunaidz/gin-rest-api/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   err.Error(),
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save user.",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully.",
		"user":    user,
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"error":   err.Error(),
		})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user.",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful.",
		"user": gin.H{
			"email":    user.Email,
			"password": user.Password,
		},
	})
}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch users.",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, users)
}
