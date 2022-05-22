package domain

type Todo struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	IsEnded bool   `json:"isEnded"`
}

type Todos []Todo

type TodoInteractor interface {
	TodosAll() (Todos, error)
}
