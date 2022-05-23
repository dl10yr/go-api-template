package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dl10yr/go-api-template/internal/domain"
	"github.com/dl10yr/go-api-template/internal/interfaces/database"
	"github.com/dl10yr/go-api-template/internal/usecase"
)

type TodoController struct {
	interactor domain.TodoInteractor
}

type Todo struct {
	Title   string `json:"name"`
	IsEnded bool   `json:"isEnded"`
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		interactor: usecase.NewTodoInteractor(
			&database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		),
	}
}

func (co *TodoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := co.interactor.TodosAll()
	if err != nil {
		panic(err)
	}

	output, _ := json.MarshalIndent(todos, "", "\t\t")

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func (co *TodoController) InsertTodo(w http.ResponseWriter, r *http.Request) {
	var in domain.TodoInput
	json.NewDecoder(r.Body).Decode(&in)
	inserted, err := co.interactor.InsertTodo(in)

	if err != nil {
		panic(err)
	}

	w.Header().Set("content-type", "application/json")
	output, _ := json.MarshalIndent(map[string]interface{}{
		"data": inserted,
	}, "", "\t\t")
	w.Write(output)
}
