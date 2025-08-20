package handlers

import (
	"class_proj_secb/models"
	"class_proj_secb/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetAllUsers(ctx *gin.Context){
	myContext := ctx.Request.Context()
	users, err := repositories.GetAllUsers(myContext)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "users fetched successfully",
		"users": users,
	})
}

func HandleCreateUser(ctx *gin.Context){
	var userDto models.UserDto
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	myContext := ctx.Request.Context()
	createdUser, err := repositories.CreateUser(myContext, userDto)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"created_user": createdUser,
	})
}