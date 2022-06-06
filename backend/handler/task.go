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
		responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if len(tasks) == 0 {
		responseError(c, http.StatusNotFound, "Data Not Found")
		return
	}

	responseOK(c, http.StatusOK, tasks)
}

func (h *taskHandler) GetTask(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	task, err := h.taskPostgres.FindByID(id)

	if err != nil {
		responseError(c, http.StatusBadRequest, err.Error())
		return
	}

	if task.ID == 0 {
		responseError(c, http.StatusNotFound, "Data Not Found")
		return
	}

	responseOK(c, http.StatusOK, task)
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

		responseError(c, http.StatusBadRequest, errorMessages)
		return
	}

	stat, err := h.taskPostgres.Create(taskAddRequest)

	if err != nil {
		responseError(c, http.StatusBadRequest, err)
		return
	}

	if stat == false {
		responseError(c, http.StatusNotFound, "Data Not Found")
		return
	}

	responseOK(c, http.StatusCreated, "success")
}

func (h *taskHandler) UpdateTask(c *gin.Context) {
	var taskUpdateRequest schema.TaskUpdateRequest

	err := c.ShouldBindJSON(&taskUpdateRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, conditional: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		responseError(c, http.StatusBadRequest, err)
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	stat, err := h.taskPostgres.Update(id, taskUpdateRequest)

	if err != nil {
		responseError(c, http.StatusBadRequest, err)
		return
	}

	if stat == false {
		responseError(c, http.StatusNotFound, "Data Not Found")
		return
	}

	responseOK(c, http.StatusOK, "success")
}

func (h *taskHandler) DeleteTask(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	stat, err := h.taskPostgres.Delete(id)
	if err != nil {
		responseError(c, http.StatusBadRequest, err)
		return
	}

	if stat == false {
		responseError(c, http.StatusNotFound, "Data Not Found")
		return
	}

	responseOK(c, http.StatusOK, "success")
}

func responseOK(c *gin.Context, code int, body interface{}) {
	c.JSON(code, gin.H{
		"data": body,
	})
}

func responseError(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{
		"message":       "Failed",
		"error_key":     code,
		"error_message": message,
		"error_data":    "{}",
	})
}
