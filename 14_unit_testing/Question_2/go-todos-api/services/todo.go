package services

import (
	"go-todos-api/models"
	"go-todos-api/repositories"
)

type TodoService struct {
	repository repositories.TodoRepository
}

func InitTodoService() TodoService {
	return TodoService{
		repository: &repositories.TodoRepositoryImpl{},
	}
}

func (ts *TodoService) GetAll() ([]models.Todo, error) {
	return ts.repository.GetAll()
}

func (ts *TodoService) GetByID(id string) (models.Todo, error) {
	return ts.repository.GetByID(id)
}

func (ts *TodoService) Create(todoInput models.TodoInput) (models.Todo, error) {
	return ts.repository.Create(todoInput)
}

func (ts *TodoService) Update(todoInput models.TodoInput, id string) (models.Todo, error) {
	return ts.repository.Update(todoInput, id)
}

func (ts *TodoService) Delete(id string) error {
	return ts.repository.Delete(id)
}
