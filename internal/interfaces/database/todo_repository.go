package database

import "github.com/dl10yr/go-api-template/internal/domain"

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
		err := rows.Scan(&t.Id)
		if err != nil {
			return todos, err
		}
		todos = append(todos, t)
	}
	return
}
