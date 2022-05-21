package main

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/infrastructure/handlers"
)

func main() {

	router := http.NewServeMux()

	// todoController := controllers.NewTodoController()
	todoHandler := handlers.NewToDoHandler()
	router.HandleFunc("/todo", todoHandler.HandleRequest)

	http.ListenAndServe(":8080", router)

}
