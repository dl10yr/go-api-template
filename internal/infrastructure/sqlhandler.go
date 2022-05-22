package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/dl10yr/go-api-template/config"
	"github.com/dl10yr/go-api-template/internal/interfaces/database"
	_ "github.com/go-sql-driver/mysql"
)

type SqlHandler struct {
	Db *sql.DB
}

type SqlRows struct {
	Rows *sql.Rows
}

func NewSqlHandler() *SqlHandler {
	dataSource := fmt.Sprintf("%s:%s%s/%s", config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlDataBase)
	Db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Db = Db

	return sqlHandler
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Db.Query(statement, args...)
	if err != nil {
		return new(SqlRows).Rows, err
	}
	row := new(SqlRows)
	row.Rows = rows
	return rows, nil
}

func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}
