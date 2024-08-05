package storage

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

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
