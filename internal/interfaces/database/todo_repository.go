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

func (repo *TodoRepository) Delete(id int) (affected int, err error) {
	exe, err := repo.Execute("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		log.Print(err)
		return
	}

	rawAffected, err := exe.RowsAffected()
	if err != nil {
		return affected, err
	}
	return int(rawAffected), nil
}

func (repo *TodoRepository) Update(id int, input domain.TodoInput) (affected int, err error) {
	exe, err := repo.Execute("UPDATE todo SET title = ?, is_ended = ? WHERE id = ?", input.Title, input.IsEnded, id)
	if err != nil {
		log.Print(err)
		return
	}

	rawAffected, err := exe.RowsAffected()
	if err != nil {
		return affected, err
	}
	return int(rawAffected), nil
}
