package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tris-tux/go-task-gin/backend/schema"
	"github.com/tris-tux/go-task-gin/backend/service"
)

type taskHandler struct {
	taskPostgres service.Task
}

func NewTask(taskPostgres service.Task) *taskHandler {
	return &taskHandler{taskPostgres}
}

func (h *taskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskPostgres.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

func (h *taskHandler) GetTask(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	task, err := h.taskPostgres.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
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

// func responseOK(w http.ResponseWriter, body interface{}) {
// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(body)
// }

// func responseError(w http.ResponseWriter, code int, message string) {
// 	w.WriteHeader(code)
// 	w.Header().Set("Content-Type", "application/json")
// 	body := map[string]string{
// 		"error": message,
// 	}
// 	json.NewEncoder(w).Encode(body)
// }
