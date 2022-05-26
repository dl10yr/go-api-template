package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dl10yr/go-api-template/internal/domain"
	"github.com/dl10yr/go-api-template/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type dummyTodoInteractor struct {
	usecase.TodoInteractor
}

func (ti *dummyTodoInteractor) TodosAll() (domain.Todos, error) {
	return domain.Todos{
		{
			Id:    "1",
			Title: "111",
		},
	}, nil
}

func (ti *dummyTodoInteractor) InsertTodo(domain.TodoInput) (int, error) {
	return 1, nil
}

func (ti *dummyTodoInteractor) DeleteTodo(todoId int) (int, error) {
	return 1, nil
}

func (ti *dummyTodoInteractor) UpdateTodo(todoId int, input domain.TodoInput) (int, error) {
	return 1, nil
}

func TestTodoController_GetAllTodos(tt *testing.T) {

	controller := TodoController{
		interactor: &dummyTodoInteractor{},
	}

	tt.Run(
		"正常系",
		func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/todo", nil)
			res := httptest.NewRecorder()

			output, _ := json.MarshalIndent(domain.Todos{
				{
					Id:    "1",
					Title: "111",
				},
			}, "", "\t\t")

			controller.GetAllTodos(res, req)
			assert.Equal(t, 200, res.Result().StatusCode)
			assert.JSONEq(t, string(output), res.Body.String())

		},
	)
}

func TestTodoController_Create(tt *testing.T) {

	controller := TodoController{
		interactor: &dummyTodoInteractor{},
	}

	tt.Run(
		"正常系",
		func(t *testing.T) {
			reqBody, _ := json.MarshalIndent(map[string]interface{}{
				"title":    "test-todo",
				"is_ended": false,
			}, "", "\t\t")

			req := httptest.NewRequest(http.MethodPost, "/todo", bytes.NewBuffer(reqBody))
			res := httptest.NewRecorder()

			output, _ := json.MarshalIndent(map[string]interface{}{
				"data": 1,
			}, "", "\t\t")

			controller.InsertTodo(res, req)
			assert.Equal(t, 200, res.Result().StatusCode)
			assert.JSONEq(t, string(output), res.Body.String())
		},
	)
}

func TestTodoController_DeleteTodo(tt *testing.T) {

	controller := TodoController{
		interactor: &dummyTodoInteractor{},
	}

	tt.Run(
		"正常系",
		func(t *testing.T) {

			req, err := http.NewRequest(http.MethodDelete, "/todo/1", nil)
			if err != nil {
				t.Fatal(err)
			}
			res := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/todo/{todoId}", controller.DeleteTodo).Methods("DELETE")
			router.ServeHTTP(res, req)

			output, _ := json.MarshalIndent(map[string]interface{}{
				"deleted": 1,
			}, "", "\t\t")

			assert.Equal(t, 200, res.Result().StatusCode)
			assert.JSONEq(t, string(output), res.Body.String())

		},
	)
}

func TestTodoController_UpdateTodo(tt *testing.T) {

	controller := TodoController{
		interactor: &dummyTodoInteractor{},
	}

	tt.Run(
		"正常系",
		func(t *testing.T) {

			reqBody, _ := json.MarshalIndent(map[string]interface{}{
				"title":    "test-todo",
				"is_ended": false,
			}, "", "\t\t")

			req, err := http.NewRequest(http.MethodPut, "/todo/1", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Fatal(err)
			}
			res := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/todo/{todoId}", controller.UpdateTodo).Methods("PUT")
			router.ServeHTTP(res, req)

			output, _ := json.MarshalIndent(map[string]interface{}{
				"updated": 1,
			}, "", "\t\t")

			assert.Equal(t, 200, res.Result().StatusCode)
			assert.JSONEq(t, string(output), res.Body.String())

		},
	)
}
