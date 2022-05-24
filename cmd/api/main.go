package main

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/infrastructure"
	"github.com/dl10yr/go-api-template/internal/infrastructure/httphandlers"
	"github.com/dl10yr/go-api-template/internal/interfaces/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	sqlHandler := infrastructure.NewSqlHandler()
	todoController := controllers.NewTodoController(sqlHandler)
	todoHandler := httphandlers.NewToDoHandler(todoController)
	router.HandleFunc("/todo", todoHandler.HandleRequest)
	router.HandleFunc("/todo/{todoId}", todoHandler.HandleRequest)

	http.ListenAndServe(":8080", router)

}
