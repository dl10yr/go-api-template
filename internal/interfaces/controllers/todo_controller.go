package controllers

import (
	"encoding/json"
	"net/http"
)

type TodoController interface {
	GetAllTodos(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
}

type Todo struct {
	Title   string `json:"name"`
	IsEnded bool   `json:"isEnded"`
}

func NewTodoController() TodoController {
	return &todoController{}
}

func (co *todoController) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todo := Todo{
		Title:   "todo1",
		IsEnded: true,
	}
	output, _ := json.MarshalIndent(todo, "", "\t\t")

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
