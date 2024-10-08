package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rest-api.com/restapi/models"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the data."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		// Check for specific error messages
		if err.Error() == "user not found" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "User not found."})
		} else if err.Error() == "invalid password" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password."})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "An error occurred."})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}
