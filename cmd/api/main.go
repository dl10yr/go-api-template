package main

import (
	"net/http"

	"github.com/dl10yr/go-api-template/internal/infrastructure"
	"github.com/dl10yr/go-api-template/internal/interfaces/controllers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	sqlHandler := infrastructure.NewSqlHandler()
	todoController := controllers.NewTodoController(sqlHandler)

	router.HandleFunc("/todo", todoController.GetAllTodos).Methods("GET")
	router.HandleFunc("/todo", todoController.InsertTodo).Methods("POST")
	router.HandleFunc("/todo/{todoId}", todoController.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{todoId}", todoController.DeleteTodo).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
