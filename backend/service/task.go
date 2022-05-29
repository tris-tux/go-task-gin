package service

import (
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type Task interface {
	Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error)
}

type task struct {
	repository database.Repository
}

func NewTask(repository database.Repository) *task {
	return &task{repository}
}

func (t *task) Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error) {
	actionTime, _ := taskAddRequest.ActionTime.Int64()

	task := schema.Task{
		Title:      taskAddRequest.Title,
		ActionTime: int(actionTime),
		CreateTime: int(actionTime),
		UpdateTime: int(actionTime),
		IsFinished: false,
	}

	newTask, err := t.repository.Create(task)

	detail := []schema.Detail{}
	details := taskAddRequest.ObjectiveList
	for _, v := range details {
		dtl := schema.Detail{
			ObjectTaskFK: newTask.ID,
			ObjectName:   v,
			IsFinished:   false,
		}
		detail = append(detail, dtl)
	}

	newDetail, err := t.repository.CreateDetail(detail)

	return newTask, newDetail, err
}
