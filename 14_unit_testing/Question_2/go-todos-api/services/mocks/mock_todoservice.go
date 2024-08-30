package mocks

import (
	"go-todos-api/models"

	"github.com/stretchr/testify/mock"
)

type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) GetAll() ([]models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]models.Todo), args.Error(1)
}

func (m *MockTodoService) GetByID(id string) (models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoService) Create(todoInput models.TodoInput) (models.Todo, error) {
	args := m.Called(todoInput)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoService) Update(todoInput models.TodoInput, id string) (models.Todo, error) {
	args := m.Called(todoInput, id)
	return args.Get(0).(models.Todo), args.Error(1)
}

func (m *MockTodoService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
