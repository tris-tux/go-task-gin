package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/handler"
	"github.com/tris-tux/go-task-gin/backend/schema"
	"github.com/tris-tux/go-task-gin/backend/service"
)

func main() {

	// dbURL := "postgres://postgres:secret@task-gin-postgres:5432/task"
	// db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&schema.Task{})
	db.AutoMigrate(&schema.Detail{})

	taskRepository := database.NewRepo(db)
	taskService := service.NewTask(taskRepository)
	taskHandler := handler.NewTask(taskService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/task/get", taskHandler.GetTasks)
	v1.GET("/task/get/:id", taskHandler.GetTask)
	v1.POST("/task/add", taskHandler.CreateTask)
	v1.PUT("/task/update/:id", taskHandler.UpdateTask)
	v1.DELETE("/task/delete/:id", taskHandler.DeleteTask)

	router.Run()
}

//main
//handler
//service
//repository
//db
//postgres
