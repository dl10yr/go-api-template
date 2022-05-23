package usecase

import "github.com/dl10yr/go-api-template/internal/domain"

type TodoInteractor struct {
	todoRepository TodoRepository
}

func NewTodoInteractor(repo TodoRepository) domain.TodoInteractor {
	return &TodoInteractor{
		todoRepository: repo,
	}
}

type TodoRepository interface {
	GetAll() (domain.Todos, error)
	Insert(todoInput domain.TodoInput) (int, error)
}

func (in *TodoInteractor) TodosAll() (todos domain.Todos, err error) {
	todos, err = in.todoRepository.GetAll()
	return
}

func (in *TodoInteractor) InsertTodo(todoInput domain.TodoInput) (int, error) {
	return in.todoRepository.Insert(todoInput)
}
