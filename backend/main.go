package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dbURL := "postgres://postgres:secret@task-gin-postgres:5432/task"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&schema.Task{})
	// db.AutoMigrate(&schema.Detail{})

	taskRepository := database.NewRepo(db)
	taskPostgres := database.NewPostgres(taskRepository)
	taskHandler := handler.NewTaskHandler(taskPostgres)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.POST("/task/add", taskHandler.CreateTask)

	router.Run()
}
