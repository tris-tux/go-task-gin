package database_test

import (
	"errors"

	"github.com/stretchr/testify/mock"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type RepositoryMock struct {
	Mock mock.Mock
}

// func (r *RepositoryMock) FindTaskAll() ([]schema.Task, error) {
// 	// arg := r.Mock.Called()
// 	// if arg != nil {
// 	// 	task := arg.Get(0).([]schema.Task)
// 	// 	return task, nil
// 	// }
// 	return nil, errors.New("Category Not Found")
// }

func (r *RepositoryMock) FindTaskByID(ID int) (*schema.Task, error) {
	arg := r.Mock.Called(ID)
	if arg.Get(0) != nil {
		task := arg.Get(0).(schema.Task)
		return &task, nil
	}
	return nil, errors.New("Category Not Found")
}

func (r *RepositoryMock) FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error) {
	arg := r.Mock.Called(ObjectTaskFK)
	if arg.Get(0) != nil {
		detail := arg.Get(0).([]schema.Detail)
		return detail, nil
	}
	return nil, errors.New("Details Not Found")
}
