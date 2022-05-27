package database

import "github.com/tris-tux/go-task-gin/backend/schema"

type Postgres interface {
	Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error)
}

type postgres struct {
	repository Repository
}

func NewPostgres(repository Repository) *postgres {
	return &postgres{repository}
}

func (p *postgres) Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error) {
	actionTime, _ := taskAddRequest.ActionTime.Int64()

	task := schema.Task{
		Title:      taskAddRequest.Title,
		ActionTime: int(actionTime),
		CreateTime: int(actionTime),
		UpdateTime: int(actionTime),
		IsFinished: false,
	}

	newTask, err := p.repository.Create(task)

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

	newDetail, err := p.repository.CreateDetail(detail)

	return newTask, newDetail, err
}
