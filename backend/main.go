package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/handler"
	"github.com/tris-tux/go-task-gin/backend/middleware"
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
	db.AutoMigrate(&schema.User{})

	userRepository := database.NewUserRepository(db)
	jwtService := service.NewJWTService()
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)
	userHandler := handler.NewUserHandler(userService, jwtService)
	authHandler := handler.NewAuthHandler(authService, jwtService)

	taskRepository := database.NewRepo(db)
	taskService := service.NewTask(taskRepository)
	taskHandler := handler.NewTask(taskService)

	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	v1 := r.Group("/v1", middleware.AuthorizeJWT(jwtService))
	{
		v1.GET("/task/get", taskHandler.GetTasks)
		v1.GET("/task/get/:id", taskHandler.GetTask)
		v1.POST("/task/add", taskHandler.CreateTask)
		v1.PUT("/task/update/:id", taskHandler.UpdateTask)
		v1.DELETE("/task/delete/:id", taskHandler.DeleteTask)
	}
	r.Run()
}

//main
//handler
//service
//repository
//db
//postgres
