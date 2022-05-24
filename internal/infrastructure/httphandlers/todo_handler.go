package httphandlers

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/interfaces/controllers"
	"github.com/gorilla/mux"
)

type Handler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	todoController *controllers.TodoController
}

func NewToDoHandler(todoController *controllers.TodoController) Handler {
	return &handler{todoController}
}

func (ha *handler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["todoId"] != "" {
		switch r.Method {
		case http.MethodDelete:
			ha.todoController.DeleteTodo(w, r)
		case http.MethodPut:
			ha.todoController.UpdateTodo(w, r)
		}
		return
	}
	switch r.Method {
	case http.MethodGet:
		ha.todoController.GetAllTodos(w, r)
	case http.MethodPost:
		ha.todoController.InsertTodo(w, r)
	}
}
