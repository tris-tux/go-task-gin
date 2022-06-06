package database

import (
	"github.com/tris-tux/go-task-gin/backend/schema"
	"gorm.io/gorm"
)

type Repository interface {
	FindTaskAll() ([]schema.Task, error)
	FindTaskByID(ID int) (schema.Task, error)
	FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error)
	Create(task schema.Task) (int, error)
	CreateDetail(detail []schema.Detail) (bool, error)
	UpdateTask(task schema.Task) (bool, error)
	DeleteTask(task schema.Task) (bool, error)
	DeleteDetails(ID int) (bool, error)
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

	return tasks, err
}

func (r *repository) FindTaskByID(ID int) (schema.Task, error) {
	var task schema.Task
	err := r.db.Find(&task, ID).Error

	return task, err
}

func (r *repository) FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error) {
	var details []schema.Detail
	err := r.db.Where("object_task_fk = ?", ObjectTaskFK).Find(&details).Error

	return details, err
}

func (r *repository) Create(task schema.Task) (int, error) {
	err := r.db.Create(&task).Error

	return task.ID, err
}

func (r *repository) CreateDetail(detail []schema.Detail) (bool, error) {
	err := r.db.Create(&detail).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) UpdateTask(task schema.Task) (bool, error) {
	err := r.db.Save(&task).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *repository) DeleteTask(task schema.Task) (bool, error) {
	err := r.db.Delete(&task).Error

	if err != nil {
		return false, err
	}

	return true, err
}

func (r *repository) DeleteDetails(ID int) (bool, error) {
	var details []schema.Detail
	err := r.db.Where("object_task_fk = ?", ID).Delete(&details).Error

	if err != nil {
		return false, err
	}

	return true, err
}
