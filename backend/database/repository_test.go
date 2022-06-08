package database_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tris-tux/go-task-gin/backend/schema"
)

type RepositoryMock struct {
	Mock mock.Mock
}

var testRepository = &RepositoryMock{Mock: mock.Mock{}}

// func TestFindTaskAlla(t *testing.T) {
// 	t.Run("FindTaskAll Data Not Found", func(t *testing.T) {
// 		testRepository.Mock.On("FindTaskAll").Return(nil)
// 		task, err := testRepository.FindTaskAll()
// 		assert.Nil(t, task)
// 		assert.NotNil(t, err)
// 	})
// 	// t.Run("FindTaskAll Data Success", func(t *testing.T) {
// 	// 	task := []schema.Task{
// 	// 		{
// 	// 			ID:         2,
// 	// 			Title:      "Tris",
// 	// 			ActionTime: 1,
// 	// 			CreateTime: 22,
// 	// 			UpdateTime: 333,
// 	// 			IsFinished: false,
// 	// 		},
// 	// 		{
// 	// 			ID:         3,
// 	// 			Title:      "Anto",
// 	// 			ActionTime: 333,
// 	// 			CreateTime: 22,
// 	// 			UpdateTime: 111,
// 	// 			IsFinished: true,
// 	// 		},
// 	// 	}
// 	// 	testRepository.Mock.On("FindTaskAll").Return(task)
// 	// 	result, err := testRepository.FindTaskAll()
// 	// 	assert.Nil(t, err)
// 	// 	assert.NotNil(t, result)
// 	// 	assert.Equal(t, task, result)
// 	// })
// }

func (r *RepositoryMock) FindTaskByID(ID int) (*schema.Task, error) {
	arg := r.Mock.Called(ID)
	if arg.Get(0) != nil {
		task := arg.Get(0).(schema.Task)
		return &task, nil
	}
	return nil, errors.New("Category Not Found")
}

func TestFindTaskByID(t *testing.T) {
	t.Run("FindTaskByID Data Not Found", func(t *testing.T) {
		testRepository.Mock.On("FindTaskByID", 1).Return(nil)
		task, err := testRepository.FindTaskByID(1)
		assert.Nil(t, task)
		assert.NotNil(t, err)
	})
	t.Run("FindTaskByID Get Data Success", func(t *testing.T) {
		task := schema.Task{
			ID:         2,
			Title:      "Tris",
			ActionTime: 1,
			CreateTime: 22,
			UpdateTime: 333,
			IsFinished: false,
		}
		testRepository.Mock.On("FindTaskByID", 2).Return(task)
		result, err := testRepository.FindTaskByID(2)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &task, result)
	})
}

func (r *RepositoryMock) FindDetailByObjectTaskFK(ObjectTaskFK int) ([]schema.Detail, error) {
	arg := r.Mock.Called(ObjectTaskFK)
	if arg.Get(0) != nil {
		detail := arg.Get(0).([]schema.Detail)
		return detail, nil
	}
	return nil, errors.New("Details Not Found")
}

func TestFindDetailByObjectTaskFK(t *testing.T) {
	t.Run("FindDetailByObjectTaskFK Data Not Found", func(t *testing.T) {
		testRepository.Mock.On("FindDetailByObjectTaskFK", 1).Return(nil)
		detail, err := testRepository.FindDetailByObjectTaskFK(1)
		assert.Nil(t, detail)
		assert.NotNil(t, err)
	})
	t.Run("FindDetailByObjectTaskFK Get Data Success", func(t *testing.T) {
		detail := []schema.Detail{
			{
				ID:           1,
				ObjectTaskFK: 2,
				ObjectName:   "Sub 1",
				IsFinished:   true,
			},
			{
				ID:           2,
				ObjectTaskFK: 2,
				ObjectName:   "Sub 2",
				IsFinished:   false,
			},
		}
		testRepository.Mock.On("FindDetailByObjectTaskFK", 2).Return(detail)
		result, err := testRepository.FindDetailByObjectTaskFK(2)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, detail, result)
	})
}
