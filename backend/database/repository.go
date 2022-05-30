package database

import (
	"github.com/tris-tux/go-task-gin/backend/schema"
	"gorm.io/gorm"
)

type Repository interface {
	FindTaskAll() ([]schema.Task, error)
	FindTaskByID(ID int) (schema.Task, error)
	FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error)
	Create(task schema.Task) (schema.Task, error)
	CreateDetail(detail []schema.Detail) ([]schema.Detail, error)
	UpdateTask(task schema.Task) (schema.Task, error)
	DeleteTask(task schema.Task) (schema.Task, error)
	DeleteDetails(ID int) (int, error)
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

func (r *repository) Create(task schema.Task) (schema.Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}

func (r *repository) CreateDetail(detail []schema.Detail) ([]schema.Detail, error) {
	err := r.db.Create(&detail).Error

	return detail, err
}

func (r *repository) UpdateTask(task schema.Task) (schema.Task, error) {
	err := r.db.Save(&task).Error

	return task, err
}

func (r *repository) DeleteTask(task schema.Task) (schema.Task, error) {
	err := r.db.Delete(&task).Error

	return task, err
}

func (r *repository) DeleteDetails(ID int) (int, error) {
	var details []schema.Detail
	err := r.db.Where("object_task_fk = ?", ID).Delete(&details).Error

	return ID, err
}
