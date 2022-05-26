package usecase

import (
	"testing"

	"github.com/dl10yr/go-api-template/internal/domain"
	"github.com/dl10yr/go-api-template/internal/interfaces/database"
	"github.com/stretchr/testify/assert"
)

type todoRepository struct {
	database.TodoRepository
}

func (repo *todoRepository) GetAll() (todos domain.Todos, err error) {
	return domain.Todos{}, nil
}

func (repo *todoRepository) Insert(input domain.TodoInput) (int, error) {
	return 1, nil
}

func (repo *todoRepository) Delete(todoId int) (int, error) {
	return 1, nil
}

func (repo *todoRepository) Update(todoId int, input domain.TodoInput) (int, error) {
	return 1, nil
}

func TestTodoInteractor_GetAll(tt *testing.T) {

	mockReviewRepo := new(todoRepository)
	interactor := NewTodoInteractor(mockReviewRepo)

	tt.Run(
		"正常系",
		func(t *testing.T) {

			r, err := interactor.TodosAll()
			assert.Equal(t, r, domain.Todos{})
			assert.Equal(t, err, nil)
		},
	)
}

func TestTodoInteractor_InsertTodo(tt *testing.T) {

	mockReviewRepo := new(todoRepository)
	interactor := NewTodoInteractor(mockReviewRepo)

	tt.Run(
		"正常系",
		func(t *testing.T) {
			todoInput := domain.TodoInput{
				Title:   "test-todo",
				IsEnded: false,
			}

			r, err := interactor.InsertTodo(todoInput)
			assert.Equal(t, r, 1)
			assert.Equal(t, err, nil)
		},
	)
}

func TestTodoInteractor_DeleteTodo(tt *testing.T) {

	mockReviewRepo := new(todoRepository)
	interactor := NewTodoInteractor(mockReviewRepo)

	tt.Run(
		"正常系",
		func(t *testing.T) {
			todoId := 1

			r, err := interactor.DeleteTodo(todoId)
			assert.Equal(t, r, 1)
			assert.Equal(t, err, nil)
		},
	)
}

func TestTodoInteractor_UpdateTodo(tt *testing.T) {

	mockReviewRepo := new(todoRepository)
	interactor := NewTodoInteractor(mockReviewRepo)

	tt.Run(
		"正常系",
		func(t *testing.T) {
			todoId := 1
			todoInput := domain.TodoInput{
				Title:   "test-todo",
				IsEnded: false,
			}

			r, err := interactor.UpdateTodo(todoId, todoInput)
			assert.Equal(t, r, 1)
			assert.Equal(t, err, nil)
		},
	)
}
