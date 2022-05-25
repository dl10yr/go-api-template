package database_test

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dl10yr/go-api-template/internal/domain"
	"github.com/dl10yr/go-api-template/internal/infrastructure"
	"github.com/dl10yr/go-api-template/internal/interfaces/database"
	"github.com/stretchr/testify/assert"
)

func TestTodoRepository_Create(tt *testing.T) {
	tt.Run(
		"正常系",
		func(t *testing.T) {
			todoInput := domain.TodoInput{
				Title:   "test-todo",
				IsEnded: false,
			}

			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO todo`)).
				ExpectExec().
				WithArgs(todoInput.Title, todoInput.IsEnded).
				WillReturnResult(sqlmock.NewResult(1, 1))

			repo := &database.TodoRepository{
				SqlHandler: infrastructure.NewDummyHandler(db),
			}

			_, err = repo.Insert(todoInput)

			assert.NoError(t, err)

			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestTodoRepository_GetAllTodos(tt *testing.T) {
	tt.Run(
		"正常系",
		func(t *testing.T) {
			todo := domain.Todo{
				Id:      "1",
				Title:   "test-todo",
				IsEnded: false,
			}

			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, title FROM todo`)).
				WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(todo.Id, todo.Title))

			repo := &database.TodoRepository{
				SqlHandler: infrastructure.NewDummyHandler(db),
			}

			ret, err := repo.GetAll()

			assert.NoError(t, err)
			assert.Equal(t, ret[0].Id, todo.Id)
			assert.Equal(t, ret[0].Title, todo.Title)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestTodoRepository_DeleteTodo(tt *testing.T) {
	tt.Run(
		"正常系",
		func(t *testing.T) {
			todoId := 1

			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM todo WHERE id = ?")).
				ExpectExec().
				WithArgs(todoId).
				WillReturnResult(sqlmock.NewResult(1, 1))

			repo := &database.TodoRepository{
				SqlHandler: infrastructure.NewDummyHandler(db),
			}
			_, err = repo.Delete(todoId)

			assert.NoError(t, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}

func TestTodoRepository_UpdateTodo(tt *testing.T) {
	tt.Run(
		"正常系",
		func(t *testing.T) {
			todoId := 1
			todoInput := domain.TodoInput{
				Title:   "test-todo",
				IsEnded: true,
			}

			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			mock.ExpectPrepare(regexp.QuoteMeta("UPDATE todo SET title = ?, is_ended = ? WHERE id = ?")).
				ExpectExec().
				WithArgs(todoInput.Title, todoInput.IsEnded, todoId).
				WillReturnResult(sqlmock.NewResult(1, 1))

			repo := &database.TodoRepository{
				SqlHandler: infrastructure.NewDummyHandler(db),
			}

			_, err = repo.Update(todoId, todoInput)

			assert.NoError(t, err)

			assert.NoError(t, mock.ExpectationsWereMet())
		},
	)
}
