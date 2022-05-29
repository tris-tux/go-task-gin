package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	dbURL, err := loadPostgresConfig()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadPostgresConfig() (string, error) {
	if os.Getenv("DB_HOST") == "" {
		return "", fmt.Errorf("Environment variable DB_HOST must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		return "", fmt.Errorf("Environment variable DB_PORT must be set")
	}
	if os.Getenv("DB_USER") == "" {
		return "", fmt.Errorf("Environment variable DB_USER must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return "", fmt.Errorf("Environment variable DB_DATABASE must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return "", fmt.Errorf("Environment variable DB_PASSWORD must be set")
	}
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	return dbURL, nil
}

// type Postgres interface {
// 	Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error)
// }

// type postgre struct {
// 	repository Repository
// }

// func NewPostgres(repository Repository) *postgre {
// 	return &postgre{repository}
// }

// func (p *postgre) Create(taskAddRequest schema.TaskAddRequest) (schema.Task, []schema.Detail, error) {
// 	actionTime, _ := taskAddRequest.ActionTime.Int64()

// 	task := schema.Task{
// 		Title:      taskAddRequest.Title,
// 		ActionTime: int(actionTime),
// 		CreateTime: int(actionTime),
// 		UpdateTime: int(actionTime),
// 		IsFinished: false,
// 	}

// 	newTask, err := p.repository.Create(task)

// 	detail := []schema.Detail{}
// 	details := taskAddRequest.ObjectiveList
// 	for _, v := range details {
// 		dtl := schema.Detail{
// 			ObjectTaskFK: newTask.ID,
// 			ObjectName:   v,
// 			IsFinished:   false,
// 		}
// 		detail = append(detail, dtl)
// 	}

// 	newDetail, err := p.repository.CreateDetail(detail)

// 	return newTask, newDetail, err
// }
