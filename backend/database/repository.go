package database

import (
	"github.com/tris-tux/go-task-gin/backend/schema"
	"gorm.io/gorm"
)

type Repository interface {
	FindTaskAll() ([]schema.Task, error)
	FindTaskByID(ID int) (*schema.Task, error)
	FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error)
	Create(task schema.Task) (int, error)
	CreateDetail(detail []schema.Detail) error
	UpdateTask(task schema.Task) error
	DeleteTask(task schema.Task) error
	DeleteDetails(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTaskAll() ([]schema.Task, error) {
	var tasks []schema.Task
	err := r.db.Find(&tasks).Error

	if err != nil {
		return nil, schema.ErrorWrap(503, err.Error())
	}

	return tasks, nil
}

func (r *repository) FindTaskByID(ID int) (*schema.Task, error) {
	var task schema.Task
	err := r.db.Find(&task, ID).Error
	if task.ID == 0 {
		return nil, schema.ErrorWrap(404, "Data Not Found")
	}
	if err != nil {
		return nil, schema.ErrorWrap(503, err.Error())
	}

	return &task, nil
}

func (r *repository) FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error) {
	var details []schema.Detail
	err := r.db.Where("object_task_fk = ?", ObjectTaskFK).Find(&details).Error

	if err != nil {
		return nil, schema.ErrorWrap(503, err.Error())
	}

	return details, nil
}

func (r *repository) Create(task schema.Task) (int, error) {
	err := r.db.Create(&task).Error

	return task.ID, err
}

func (r *repository) CreateDetail(detail []schema.Detail) error {
	err := r.db.Create(&detail).Error

	return err
}

func (r *repository) UpdateTask(task schema.Task) error {
	err := r.db.Save(&task).Error

	return err
}

func (r *repository) DeleteTask(task schema.Task) error {
	err := r.db.Delete(&task).Error

	return err
}

func (r *repository) DeleteDetails(ID int) error {
	var details []schema.Detail
	err := r.db.Where("object_task_fk = ?", ID).Delete(&details).Error

	return err
}
