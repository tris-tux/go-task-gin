package service

import (
	"reflect"
	"testing"

	"github.com/tris-tux/go-task-gin/backend/database"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

// var taskRepository = &database.RepositoryMock{Mock: mock.Mock{}}
// var taskService = task{repository: taskRepository}

// func TestTaskGetNotFound(t *testing.T) {
// 	taskRepository.Mock.On("FindID", 1).Return(nil)
// 	task, err := taskService.FindID(1)
// 	assert.Nil(t, task)
// 	assert.NotNil(t, err)
// }

// func TestTaskGetGetSuccess(t *testing.T) {
// 	task := schema.Task{
// 		ID:         2,
// 		Title:      "Anto",
// 		ActionTime: 1234324,
// 		CreateTime: 2143234,
// 		UpdateTime: 2343243,
// 		IsFinished: true,
// 	}

// 	taskRepository.Mock.On("FindID", 2).Return(task)

// 	result, err := taskService.FindID(2)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, task.ID, result.ID)
// 	assert.Equal(t, task.Title, result.Title)
// }

// func TestFindByIDGetNotFound(t *testing.T) {
// 	taskRepository.Mock.On("FindByID", 1).Return(nil)
// 	task, err := taskService.FindByID(1)
// 	assert.Nil(t, task)
// 	assert.NotNil(t, err)
// }

func Test_task_FindByID(t *testing.T) {
	type fields struct {
		repository database.Repository
	}
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    schema.TaskResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &task{
				repository: tt.fields.repository,
			}
			got, err := tr.FindByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("task.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("task.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
