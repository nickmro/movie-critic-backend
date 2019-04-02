package postgres

import "database/sql"

// DB represents a PostgreSQL DB connection.
type DB struct {
	*sql.DB
}

// MustBegin begins a database transaction or panics.
func (db *DB) MustBegin() *Tx {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	return &Tx{tx}
}

// Query queries the database.
func (db *DB) Query(query string, args []interface{}) (*sql.Rows, error) {
	return db.DB.Query(query, args...)
}
