package service

import (
	"time"

	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type Task interface {
	FindAll() ([]schema.TaskResponse, error)
	FindByID(ID int) (schema.TaskResponse, error)
	Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error)
	Update(ID int, taskUpdateRequest schema.TaskUpdateRequest) (schema.Task, []schema.Detail, error)
	Delete(ID int) (schema.Task, int, error)
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

	newTask := schema.TaskResponse{
		ID:         task.ID,
		Title:      task.Title,
		ActionTime: task.ActionTime,
		CreateTime: task.CreateTime,
		UpdateTime: task.UpdateTime,
		IsFinished: task.IsFinished,
		ObjectList: detailAll,
	}

	return newTask, err
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

func (t *task) Update(ID int, taskUpdateRequest schema.TaskUpdateRequest) (schema.Task, []schema.Detail, error) {
	task, err := t.repository.FindTaskByID(ID)

	now := time.Now()
	timeStamp := now.Unix()

	task.Title = taskUpdateRequest.Title
	task.CreateTime = task.ActionTime
	task.UpdateTime = int(timeStamp)
	task.ActionTime = int(timeStamp)
	task.IsFinished = false

	newTask, err := t.repository.UpdateTask(task)

	newDelete, err := t.repository.DeleteDetails(ID)

	detail := []schema.Detail{}
	details := taskUpdateRequest.ObjectiveList
	for _, v := range details {
		dtl := schema.Detail{
			ObjectTaskFK: newDelete,
			ObjectName:   v.ObjectName,
			IsFinished:   v.IsFinished,
		}
		detail = append(detail, dtl)
	}

	newDetail, err := t.repository.CreateDetail(detail)
	// task, err := t.repository.UpdateTask(task)

	return newTask, newDetail, err
}

func (t *task) Delete(ID int) (schema.Task, int, error) {
	task, err := t.repository.FindTaskByID(ID)

	newTask, err := t.repository.DeleteTask(task)
	newDelete, err := t.repository.DeleteDetails(ID)
	return newTask, newDelete, err
}
