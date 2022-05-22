package httphandlers

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/interfaces/controllers"
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
	switch r.Method {
	case http.MethodGet:
		ha.todoController.GetAllTodos(w, r)
	}
}
