package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type taskHandler struct {
	taskPostgres database.Postgres
}

func NewTaskHandler(taskPostgres database.Postgres) *taskHandler {
	return &taskHandler{taskPostgres}
}

func (h *taskHandler) CreateTask(c *gin.Context) {
	var taskAddRequest schema.TaskAddRequest

	err := c.ShouldBindJSON(&taskAddRequest)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, conditional: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	task, detail, err := h.taskPostgres.Create(taskAddRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task":   task,
		"detail": detail,
	})
}
