package database

import (
	"log"

	"github.com/dl10yr/go-api-template/internal/domain"
)

type TodoRepository struct {
	SqlHandler
}

func (repo *TodoRepository) GetAll() (todos domain.Todos, err error) {
	rows, err := repo.Query("SELECT id, title FROM todo")

	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var t domain.Todo
		err := rows.Scan(&t.Id, &t.Title)
		if err != nil {
			return todos, err
		}
		todos = append(todos, t)
	}
	return
}

func (repo *TodoRepository) Insert(input domain.TodoInput) (inserted int, err error) {
	exe, err := repo.Execute("INSERT INTO todo (title, is_ended) VALUES (?, ?)", input.Title, input.IsEnded)
	if err != nil {
		log.Print(err)
		return
	}
	i, err := exe.LastInsertId()
	if err != nil {
		log.Print(err)
		return
	}
	return int(i), nil
}
