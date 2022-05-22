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
}

func (in *TodoInteractor) TodosAll() (todos domain.Todos, err error) {
	todos, err = in.todoRepository.GetAll()
	return
}
