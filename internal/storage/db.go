package storage

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Verify the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Successfully connected to the database")
    return db, nil
}
