package main

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/infrastructure/handlers"
	"github.com/dl10yr/go-api-template/internal/interfaces/controllers"
)

func main() {

	router := http.NewServeMux()

	todoController := controllers.NewTodoController()
	todoHandler := handlers.NewToDoHandler(todoController)
	router.HandleFunc("/todo", todoHandler.HandleRequest)

	http.ListenAndServe(":8080", router)

}
