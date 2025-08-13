package todo

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetAllTodos(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, TodoList)
}

func HandleAddTodo(ctx *gin.Context) {
	var todo Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	// todo : save todo here
	todo.ID = len(TodoList) + 1
	log.Println("todo : ", todo)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo added successfully",
		"added_todo": todo,
	})
}

func HandleDeleteTodo(ctx *gin.Context){
	id := ctx.Param("id")
	log.Println("id is : ", id)
	// todo logic to delete todo
	ctx.JSON(http.StatusOK, gin.H{
		"message": "todo deleted successfully",
		"deleted_id": id,
	} )
}
