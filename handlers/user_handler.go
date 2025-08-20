package handlers

import (
	"class_proj_secb/models"
	"class_proj_secb/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleGetAllUsers(ctx *gin.Context) {
	myContext := ctx.Request.Context()
	users, err := repositories.GetAllUsers(myContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "users fetched successfully",
		"users":   users,
	})
}

func HandleCreateUser(ctx *gin.Context) {
	var userDto models.UserDto
	err := ctx.ShouldBindJSON(&userDto)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	myContext := ctx.Request.Context()
	createdUser, err := repositories.CreateUser(myContext, userDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "user created successfully",
		"created_user": createdUser,
	})
}

func HandleGetUserById(ctx *gin.Context) {
	myctx := ctx.Request.Context()
	id := ctx.Param("id")
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "invalid id",
		})
		return
	}
	user, err := repositories.GetUserById(myctx, idUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user fetched successfully",
		"user":    user,
	})
}

func HandleDeleteUser(ctx *gin.Context) {
	myctx := ctx.Request.Context()
	id := ctx.Param("id")
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "invalid id",
		})
		return
	}
	err = repositories.DeleteUser(myctx, idUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"id":      id,
	})
}

func HandleUpdateUser(ctx *gin.Context) {
	myctx := ctx.Request.Context()
	id := ctx.Param("id")
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "invalid id",
		})
		return
	}
	var userDto models.UserDto
	err = ctx.ShouldBindJSON(&userDto)
	if err != nil{
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
			"message": "invalid data",
		})
		return
	}
	updatedData, err := repositories.UpdateUser(myctx,userDto, idUint64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   err.Error(),
			"message": "user not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
		"updated_data": updatedData,
	})
}
