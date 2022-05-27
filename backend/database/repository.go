package database

import (
	"github.com/tris-tux/go-task-gin/backend/schema"
	"gorm.io/gorm"
)

type Repository interface {
	Create(task schema.Task) (schema.Task, error)
	CreateDetail(detail []schema.Detail) ([]schema.Detail, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(task schema.Task) (schema.Task, error) {
	err := r.db.Create(&task).Error

	return task, err
}

func (r *repository) CreateDetail(detail []schema.Detail) ([]schema.Detail, error) {
	err := r.db.Create(&detail).Error

	return detail, err
}
