package handlers

import (
    "database/sql"
    "net/http"
)

func GetPosts(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("GetPosts handler"))
    }
}

func CreatePost(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("CreatePost handler"))
    }
}
