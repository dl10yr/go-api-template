package database

type SqlHandler interface {
	Query(string, ...interface{}) (Rows, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
