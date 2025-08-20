package main

import (
	"class_proj_secb/common"
	"class_proj_secb/database/myquery"
	"class_proj_secb/handlers"
	"class_proj_secb/models"
	"class_proj_secb/todo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	g := gin.Default()
	g.GET("/", HandleInitialRoute)
	g.GET("/todos", todo.HandleGetAllTodos)
	g.POST("/todo", todo.HandleAddTodo)
	g.DELETE("/todo/:id", todo.HandleDeleteTodo)

	// user route
	g.GET("/users", handlers.HandleGetAllUsers)
	g.POST("/create-user",handlers.HandleCreateUser)
	g.GET("/user/:id", handlers.HandleGetUserById)
	g.DELETE("/user/:id", handlers.HandleDeleteUser)
	g.PUT("/update-user/:id", handlers.HandleUpdateUser)

	// checking database connection
	dsn := "host=localhost user=secb password=secb dbname=secbgodb port=5434 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Println("failed to connect database")
		return;
	}
	// using db in myquery
	myquery.SetDefault(db)
	// migrating user table
	db.AutoMigrate(models.User{})
	log.Println(db)
	if db != nil{
		common.MyApiServer.DB = *db
	}else{
		log.Fatal("Database not connected")
		return
	}
	log.Println("database connected :5434")
	// running and listening on port
	// default port :8080
	g.Run(":8010")
}

func HandleInitialRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Welocome to the golang api",
	})
}
