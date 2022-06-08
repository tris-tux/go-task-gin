package service

// import (
// 	"errors"

// 	"github.com/stretchr/testify/mock"
// 	"github.com/tris-tux/go-task-gin/backend/schema"
// )

// type ServiceMock struct {
// 	Mock mock.Mock
// }

// func (r *ServiceMock) FindByID(ID int) (*schema.Task, error) {
// 	arg := r.Mock.Called(ID)
// 	if arg.Get(0) != nil {
// 		task := arg.Get(0).(schema.Task)
// 		return &task, nil
// 	}
// 	return nil, errors.New("Category Not Found")
// }
