package service

import (
	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type Task interface {
	FindAll() ([]schema.TaskResponse, error)
	FindByID(ID int) (schema.TaskResponse, error)
	Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error)
}

type task struct {
	repository database.Repository
}

func NewTask(repository database.Repository) *task {
	return &task{repository}
}

func (t *task) FindAll() ([]schema.TaskResponse, error) {
	taskAll := []schema.TaskResponse{}
	tasks, err := t.repository.FindTaskAll()
	for _, v := range tasks {
		details, _ := t.repository.FindDetailByObjectTaskFK(v.ID)
		detailAll := []schema.DetailResponse{}
		for _, d := range details {
			detail := schema.DetailResponse{
				ObjectName: d.ObjectName,
				IsFinished: d.IsFinished,
			}

			detailAll = append(detailAll, detail)
		}

		task := schema.TaskResponse{
			ID:         v.ID,
			Title:      v.Title,
			ActionTime: v.ActionTime,
			CreateTime: v.CreateTime,
			UpdateTime: v.UpdateTime,
			IsFinished: v.IsFinished,
			ObjectList: detailAll,
		}
		taskAll = append(taskAll, task)

	}
	return taskAll, err
}

func (t *task) FindByID(ID int) (schema.TaskResponse, error) {
	task, err := t.repository.FindTaskByID(ID)

	details, _ := t.repository.FindDetailByObjectTaskFK(task.ID)
	detailAll := []schema.DetailResponse{}
	for _, d := range details {
		detail := schema.DetailResponse{
			ObjectName: d.ObjectName,
			IsFinished: d.IsFinished,
		}

		detailAll = append(detailAll, detail)
	}

	taskRes := schema.TaskResponse{
		ID:         task.ID,
		Title:      task.Title,
		ActionTime: task.ActionTime,
		CreateTime: task.CreateTime,
		UpdateTime: task.UpdateTime,
		IsFinished: task.IsFinished,
		ObjectList: detailAll,
	}

	return taskRes, err
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
