package domain

type Todo struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	IsEnded bool   `json:"is_ended"`
}

type Todos []Todo

type TodoInteractor interface {
	TodosAll() (Todos, error)
	InsertTodo(input TodoInput) (int, error)
	DeleteTodo(todoId int) (int, error)
	// PutTodo(input TodoInput) (int, error)
}

type TodoInput struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	IsEnded bool   `json:"is_ended"`
}
