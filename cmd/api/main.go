package main

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/infrastructure"
	"github.com/dl10yr/go-api-template/internal/infrastructure/httphandlers"
	"github.com/dl10yr/go-api-template/internal/interfaces/controllers"
)

func main() {

	router := http.NewServeMux()

	sqlHandler := infrastructure.NewSqlHandler()
	todoController := controllers.NewTodoController(sqlHandler)
	todoHandler := httphandlers.NewToDoHandler(todoController)
	router.HandleFunc("/todo", todoHandler.HandleRequest)

	http.ListenAndServe(":8080", router)

}
