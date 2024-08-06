package storage

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

// Represents DB configuration
//
// Fields:
// - *sql.DB: The DB connection
//
// Since: 0.0.1
type DB struct {
    *sql.DB
}

// Create new DB connection
//
// Params:
// - dataSourceName (string): the DB URL
//
// Returns:
// - *DB, error: If successful, returns the database connection
// otherwise return the error
//
// Since: 0.0.1
func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Verify the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Successfully connected to the database")
    return &DB{db}, nil
}
