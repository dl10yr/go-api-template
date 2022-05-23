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

type SqlResult struct {
	Result sql.Result
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

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	stmt, err := handler.Db.Prepare(statement)
	if err != nil {
		return res.Result, err
	}
	defer stmt.Close()
	exe, err := stmt.Exec(args...)
	if err != nil {
		return res.Result, err
	}
	res.Result = exe
	return res.Result, nil
}

func (r SqlRows) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}
