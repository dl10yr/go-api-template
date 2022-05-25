package infrastructure

import (
	"database/sql"

	"github.com/dl10yr/go-api-template/internal/interfaces/database"
)

func NewDummyHandler(db *sql.DB) database.SqlHandler {
	sqlHandler := new(SqlHandler)
	sqlHandler.Db = db
	return sqlHandler
}
