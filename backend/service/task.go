package service

import (
	"time"

	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type Task interface {
	FindAll() ([]schema.TaskResponse, error)
	FindByID(ID int) (*schema.TaskResponse, error)
	Create(taskAddRequest schema.TaskAddRequest) error
	Update(ID int, taskUpdateRequest schema.TaskUpdateRequest) error
	Delete(ID int) error
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
	if err != nil {
		return nil, err
	}
	for _, v := range tasks {
		details, err := t.repository.FindDetailByObjectTaskFK(v.ID)
		if err != nil {
			return nil, err
		}
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

func (t *task) FindByID(ID int) (*schema.TaskResponse, error) {
	task, err := t.repository.FindTaskByID(ID)
	if err != nil {
		return nil, err
	}

	details, err := t.repository.FindDetailByObjectTaskFK(task.ID)
	if err != nil {
		return nil, err
	}
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

	return &newTask, err
}

func (t *task) Create(taskAddRequest schema.TaskAddRequest) error {
	task := schema.Task{
		Title:      taskAddRequest.Title,
		ActionTime: taskAddRequest.ActionTime,
		CreateTime: taskAddRequest.ActionTime,
		UpdateTime: taskAddRequest.ActionTime,
		IsFinished: false,
	}

	newTask, err := t.repository.Create(task)
	if err != nil {
		return err
	}

	detail := []schema.Detail{}
	details := taskAddRequest.ObjectiveList
	for _, v := range details {
		dtl := schema.Detail{
			ObjectTaskFK: newTask,
			ObjectName:   v,
			IsFinished:   false,
		}
		detail = append(detail, dtl)
	}

	_ = t.repository.CreateDetail(detail)

	return nil
}

func (t *task) Update(ID int, taskUpdateRequest schema.TaskUpdateRequest) error {
	task, err := t.repository.FindTaskByID(ID)
	if err != nil {
		return err
	}

	now := time.Now()
	timeStamp := now.Unix()
	task.Title = taskUpdateRequest.Title
	task.CreateTime = task.ActionTime
	task.UpdateTime = int(timeStamp)
	task.ActionTime = int(timeStamp)
	task.IsFinished = false
	err = t.repository.UpdateTask(*task)
	if err != nil {
		return err
	}

	err = t.repository.DeleteDetails(ID)
	if err != nil {
		return err
	}

	detail := []schema.Detail{}
	details := taskUpdateRequest.ObjectiveList
	for _, v := range details {
		dtl := schema.Detail{
			ObjectTaskFK: ID,
			ObjectName:   v.ObjectName,
			IsFinished:   v.IsFinished,
		}
		detail = append(detail, dtl)
	}

	err = t.repository.CreateDetail(detail)
	if err != nil {
		return err
	}

	return err
}

func (t *task) Delete(ID int) error {
	task, err := t.repository.FindTaskByID(ID)
	if err != nil {
		return err
	}

	err = t.repository.DeleteTask(*task)
	if err != nil {
		return err
	}

	err = t.repository.DeleteDetails(ID)

	return err
}
